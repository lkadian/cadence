package sema

import (
	"github.com/dapperlabs/flow-go/language/runtime/ast"
	"github.com/dapperlabs/flow-go/language/runtime/common"
	"github.com/dapperlabs/flow-go/language/runtime/errors"
)

func (checker *Checker) VisitInterfaceDeclaration(declaration *ast.InterfaceDeclaration) ast.Repr {

	interfaceType := checker.Elaboration.InterfaceDeclarationTypes[declaration]

	checker.containerTypes[interfaceType] = true
	defer func() {
		checker.containerTypes[interfaceType] = false
	}()

	checker.checkDeclarationAccessModifier(
		declaration.Access,
		declaration.DeclarationKind(),
		declaration.StartPos,
		true,
		false,
	)

	// NOTE: functions are checked separately
	checker.checkFieldsAccessModifier(declaration.Members.Fields)

	checker.checkNestedIdentifiers(
		declaration.Members.Fields,
		declaration.Members.Functions,
		declaration.InterfaceDeclarations,
		declaration.CompositeDeclarations,
	)

	// Activate new scope for nested types

	checker.typeActivations.Enter()
	defer checker.typeActivations.Leave()

	// Declare nested types

	for name, nestedType := range interfaceType.NestedTypes {
		nestedDeclaration := checker.Elaboration.InterfaceNestedDeclarations[declaration][name]

		identifier := nestedDeclaration.DeclarationIdentifier()
		if identifier == nil {
			// It should be impossible to have a nested declaration
			// that does not have an identifier

			panic(errors.NewUnreachableError())
		}

		_, err := checker.typeActivations.DeclareType(
			*identifier,
			nestedType,
			nestedDeclaration.DeclarationKind(),
			nestedDeclaration.DeclarationAccess(),
		)
		checker.report(err)
	}

	// Check members

	members, origins := checker.membersAndOrigins(
		interfaceType,
		declaration.Members.Fields,
		declaration.Members.Functions,
		false,
	)

	interfaceType.Members = members

	interfaceType.InitializerParameterTypeAnnotations =
		checker.initializerParameterTypeAnnotations(declaration.Members.Initializers())

	checker.memberOrigins[interfaceType] = origins

	checker.checkInitializers(
		declaration.Members.Initializers(),
		declaration.Members.Fields,
		interfaceType,
		declaration.DeclarationKind(),
		declaration.Identifier.Identifier,
		interfaceType.InitializerParameterTypeAnnotations,
		ContainerKindInterface,
		nil,
	)

	checker.checkDestructors(
		declaration.Members.Destructors(),
		declaration.Members.FieldsByIdentifier(),
		interfaceType.Members,
		interfaceType,
		declaration.DeclarationKind(),
		declaration.Identifier.Identifier,
		ContainerKindInterface,
	)

	checker.checkUnknownSpecialFunctions(declaration.Members.SpecialFunctions)

	checker.checkInterfaceFunctions(
		declaration.Members.Functions,
		interfaceType,
		declaration.DeclarationKind(),
	)

	checker.checkResourceFieldNesting(
		declaration.Members.FieldsByIdentifier(),
		interfaceType.Members,
		interfaceType.CompositeKind,
	)

	checker.checkCompositeDeclarationSupport(
		declaration.CompositeKind,
		declaration.DeclarationKind(),
		declaration.Identifier,
	)

	return nil
}

func (checker *Checker) checkInterfaceFunctions(
	functions []*ast.FunctionDeclaration,
	interfaceType *InterfaceType,
	declarationKind common.DeclarationKind,
) {
	inResource := interfaceType.CompositeKind == common.CompositeKindResource

	for _, function := range functions {
		// NOTE: new activation, as function declarations
		// shouldn't be visible in other function declarations,
		// and `self` is is only visible inside function

		func() {
			checker.enterValueScope()
			defer checker.leaveValueScope(false)

			// NOTE: required for
			checker.declareSelfValue(interfaceType)

			checker.visitFunctionDeclaration(
				function,
				functionDeclarationOptions{
					mustExit:                false,
					declareFunction:         false,
					checkResourceLoss:       false,
					allowAuthAccessModifier: inResource,
				},
			)

			if function.FunctionBlock != nil {
				checker.checkInterfaceSpecialFunctionBlock(
					function.FunctionBlock,
					declarationKind,
					common.DeclarationKindFunction,
				)
			}
		}()
	}
}

func (checker *Checker) declareInterfaceDeclaration(declaration *ast.InterfaceDeclaration) *InterfaceType {

	identifier := declaration.Identifier

	// NOTE: fields and functions might already refer to interface itself.
	// insert a dummy type for now, so lookup succeeds during conversion,
	// then fix up the type reference

	interfaceType := &InterfaceType{
		Location:      checker.Location,
		Identifier:    identifier.Identifier,
		CompositeKind: declaration.CompositeKind,
		NestedTypes:   map[string]Type{},
	}

	variable, err := checker.typeActivations.DeclareType(
		identifier,
		interfaceType,
		declaration.DeclarationKind(),
		declaration.Access,
	)
	checker.report(err)
	checker.recordVariableDeclarationOccurrence(identifier.Identifier, variable)

	(func() {
		// Activate new scope for nested declarations

		checker.typeActivations.Enter()
		defer checker.typeActivations.Leave()

		checker.valueActivations.Enter()
		defer checker.valueActivations.Leave()

		// Check and declare nested types

		nestedDeclarations, nestedInterfaceTypes, nestedCompositeTypes :=
			checker.visitNestedDeclarations(
				declaration.CompositeKind,
				declaration.DeclarationKind(),
				declaration.CompositeDeclarations,
				declaration.InterfaceDeclarations,
			)

		checker.Elaboration.InterfaceNestedDeclarations[declaration] = nestedDeclarations

		for _, nestedInterfaceType := range nestedInterfaceTypes {
			interfaceType.NestedTypes[nestedInterfaceType.Identifier] = nestedInterfaceType
			nestedInterfaceType.ContainerType = interfaceType
		}

		for _, nestedCompositeType := range nestedCompositeTypes {
			interfaceType.NestedTypes[nestedCompositeType.Identifier] = nestedCompositeType
			nestedCompositeType.ContainerType = interfaceType
		}

	})()

	// NOTE: interface type's `InitializerParameterTypeAnnotations` and  `members` fields
	// are added in `VisitInterfaceDeclaration`.
	//
	// They are left out for now, as initializers, fields, and function requirements
	// could already refer to e.g. composites

	checker.Elaboration.InterfaceDeclarationTypes[declaration] = interfaceType

	return interfaceType
}

func (checker *Checker) checkInterfaceSpecialFunctionBlock(
	block *ast.FunctionBlock,
	containerKind common.DeclarationKind,
	implementedKind common.DeclarationKind,
) {

	if len(block.Statements) > 0 {
		checker.report(
			&InvalidImplementationError{
				Pos:             block.Statements[0].StartPosition(),
				ContainerKind:   containerKind,
				ImplementedKind: implementedKind,
			},
		)
	} else if len(block.PreConditions) == 0 &&
		len(block.PostConditions) == 0 {

		checker.report(
			&InvalidImplementationError{
				Pos:             block.StartPos,
				ContainerKind:   containerKind,
				ImplementedKind: implementedKind,
			},
		)
	}
}
