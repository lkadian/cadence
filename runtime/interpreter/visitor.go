package interpreter

type Visitor interface {
	VisitValue(interpreter *Interpreter, value Value)
	VisitTypeValue(interpreter *Interpreter, value TypeValue)
	VisitVoidValue(interpreter *Interpreter, value VoidValue)
	VisitBoolValue(interpreter *Interpreter, value BoolValue)
	VisitStringValue(interpreter *Interpreter, value *StringValue)
	VisitArrayValue(interpreter *Interpreter, value *ArrayValue) bool
	VisitIntValue(interpreter *Interpreter, value IntValue)
	VisitInt8Value(interpreter *Interpreter, value Int8Value)
	VisitInt16Value(interpreter *Interpreter, value Int16Value)
	VisitInt32Value(interpreter *Interpreter, value Int32Value)
	VisitInt64Value(interpreter *Interpreter, value Int64Value)
	VisitInt128Value(interpreter *Interpreter, value Int128Value)
	VisitInt256Value(interpreter *Interpreter, value Int256Value)
	VisitUIntValue(interpreter *Interpreter, value UIntValue)
	VisitUInt8Value(interpreter *Interpreter, value UInt8Value)
	VisitUInt16Value(interpreter *Interpreter, value UInt16Value)
	VisitUInt32Value(interpreter *Interpreter, value UInt32Value)
	VisitUInt64Value(interpreter *Interpreter, value UInt64Value)
	VisitUInt128Value(interpreter *Interpreter, value UInt128Value)
	VisitUInt256Value(interpreter *Interpreter, value UInt256Value)
	VisitWord8Value(interpreter *Interpreter, value Word8Value)
	VisitWord16Value(interpreter *Interpreter, value Word16Value)
	VisitWord32Value(interpreter *Interpreter, value Word32Value)
	VisitWord64Value(interpreter *Interpreter, value Word64Value)
	VisitFix64Value(interpreter *Interpreter, value Fix64Value)
	VisitUFix64Value(interpreter *Interpreter, value UFix64Value)
	VisitCompositeValue(interpreter *Interpreter, value *CompositeValue) bool
	VisitDictionaryValue(interpreter *Interpreter, value *DictionaryValue) bool
	VisitNilValue(interpreter *Interpreter, value NilValue)
	VisitSomeValue(interpreter *Interpreter, value *SomeValue) bool
	VisitStorageReferenceValue(interpreter *Interpreter, value *StorageReferenceValue)
	VisitEphemeralReferenceValue(interpreter *Interpreter, value *EphemeralReferenceValue)
	VisitAddressValue(interpreter *Interpreter, value AddressValue)
	VisitAuthAccountValue(interpreter *Interpreter, value AuthAccountValue)
	VisitPublicAccountValue(interpreter *Interpreter, value PublicAccountValue)
	VisitPathValue(interpreter *Interpreter, value PathValue)
	VisitCapabilityValue(interpreter *Interpreter, value CapabilityValue)
	VisitLinkValue(interpreter *Interpreter, value LinkValue)
	VisitInterpretedFunctionValue(interpreter *Interpreter, value InterpretedFunctionValue)
	VisitHostFunctionValue(interpreter *Interpreter, value HostFunctionValue)
	VisitBoundFunctionValue(interpreter *Interpreter, value BoundFunctionValue)
	VisitAuthAccountContractsValue(interpreter *Interpreter, value AuthAccountContractsValue)
	VisitDeployedContractValue(interpreter *Interpreter, value DeployedContractValue)
}

type EmptyVisitor struct {
	ValueVisitor                     func(interpreter *Interpreter, value Value)
	TypeValueVisitor                 func(interpreter *Interpreter, value TypeValue)
	VoidValueVisitor                 func(interpreter *Interpreter, value VoidValue)
	BoolValueVisitor                 func(interpreter *Interpreter, value BoolValue)
	StringValueVisitor               func(interpreter *Interpreter, value *StringValue)
	ArrayValueVisitor                func(interpreter *Interpreter, value *ArrayValue) bool
	IntValueVisitor                  func(interpreter *Interpreter, value IntValue)
	Int8ValueVisitor                 func(interpreter *Interpreter, value Int8Value)
	Int16ValueVisitor                func(interpreter *Interpreter, value Int16Value)
	Int32ValueVisitor                func(interpreter *Interpreter, value Int32Value)
	Int64ValueVisitor                func(interpreter *Interpreter, value Int64Value)
	Int128ValueVisitor               func(interpreter *Interpreter, value Int128Value)
	Int256ValueVisitor               func(interpreter *Interpreter, value Int256Value)
	UIntValueVisitor                 func(interpreter *Interpreter, value UIntValue)
	UInt8ValueVisitor                func(interpreter *Interpreter, value UInt8Value)
	UInt16ValueVisitor               func(interpreter *Interpreter, value UInt16Value)
	UInt32ValueVisitor               func(interpreter *Interpreter, value UInt32Value)
	UInt64ValueVisitor               func(interpreter *Interpreter, value UInt64Value)
	UInt128ValueVisitor              func(interpreter *Interpreter, value UInt128Value)
	UInt256ValueVisitor              func(interpreter *Interpreter, value UInt256Value)
	Word8ValueVisitor                func(interpreter *Interpreter, value Word8Value)
	Word16ValueVisitor               func(interpreter *Interpreter, value Word16Value)
	Word32ValueVisitor               func(interpreter *Interpreter, value Word32Value)
	Word64ValueVisitor               func(interpreter *Interpreter, value Word64Value)
	Fix64ValueVisitor                func(interpreter *Interpreter, value Fix64Value)
	UFix64ValueVisitor               func(interpreter *Interpreter, value UFix64Value)
	CompositeValueVisitor            func(interpreter *Interpreter, value *CompositeValue) bool
	DictionaryValueVisitor           func(interpreter *Interpreter, value *DictionaryValue) bool
	NilValueVisitor                  func(interpreter *Interpreter, value NilValue)
	SomeValueVisitor                 func(interpreter *Interpreter, value *SomeValue) bool
	StorageReferenceValueVisitor     func(interpreter *Interpreter, value *StorageReferenceValue)
	EphemeralReferenceValueVisitor   func(interpreter *Interpreter, value *EphemeralReferenceValue)
	AddressValueVisitor              func(interpreter *Interpreter, value AddressValue)
	AuthAccountValueVisitor          func(interpreter *Interpreter, value AuthAccountValue)
	PublicAccountValueVisitor        func(interpreter *Interpreter, value PublicAccountValue)
	PathValueVisitor                 func(interpreter *Interpreter, value PathValue)
	CapabilityValueVisitor           func(interpreter *Interpreter, value CapabilityValue)
	LinkValueVisitor                 func(interpreter *Interpreter, value LinkValue)
	InterpretedFunctionValueVisitor  func(interpreter *Interpreter, value InterpretedFunctionValue)
	HostFunctionValueVisitor         func(interpreter *Interpreter, value HostFunctionValue)
	BoundFunctionValueVisitor        func(interpreter *Interpreter, value BoundFunctionValue)
	AuthAccountContractsValueVisitor func(interpreter *Interpreter, value AuthAccountContractsValue)
	DeployedContractValueVisitor     func(interpreter *Interpreter, value DeployedContractValue)
}

func (v EmptyVisitor) VisitValue(interpreter *Interpreter, value Value) {
	if v.ValueVisitor == nil {
		return
	}
	v.ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitTypeValue(interpreter *Interpreter, value TypeValue) {
	if v.TypeValueVisitor == nil {
		return
	}
	v.TypeValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitVoidValue(interpreter *Interpreter, value VoidValue) {
	if v.VoidValueVisitor == nil {
		return
	}
	v.VoidValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitBoolValue(interpreter *Interpreter, value BoolValue) {
	if v.BoolValueVisitor == nil {
		return
	}
	v.BoolValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitStringValue(interpreter *Interpreter, value *StringValue) {
	if v.StringValueVisitor == nil {
		return
	}
	v.StringValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitArrayValue(interpreter *Interpreter, value *ArrayValue) bool {
	if v.ArrayValueVisitor == nil {
		return true
	}
	return v.ArrayValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitIntValue(interpreter *Interpreter, value IntValue) {
	if v.IntValueVisitor == nil {
		return
	}
	v.IntValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitInt8Value(interpreter *Interpreter, value Int8Value) {
	if v.Int8ValueVisitor == nil {
		return
	}
	v.Int8ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitInt16Value(interpreter *Interpreter, value Int16Value) {
	if v.Int16ValueVisitor == nil {
		return
	}
	v.Int16ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitInt32Value(interpreter *Interpreter, value Int32Value) {
	if v.Int32ValueVisitor == nil {
		return
	}
	v.Int32ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitInt64Value(interpreter *Interpreter, value Int64Value) {
	if v.Int64ValueVisitor == nil {
		return
	}
	v.Int64ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitInt128Value(interpreter *Interpreter, value Int128Value) {
	if v.Int128ValueVisitor == nil {
		return
	}
	v.Int128ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitInt256Value(interpreter *Interpreter, value Int256Value) {
	if v.Int256ValueVisitor == nil {
		return
	}
	v.Int256ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUIntValue(interpreter *Interpreter, value UIntValue) {
	if v.UIntValueVisitor == nil {
		return
	}
	v.UIntValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUInt8Value(interpreter *Interpreter, value UInt8Value) {
	if v.UInt8ValueVisitor == nil {
		return
	}
	v.UInt8ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUInt16Value(interpreter *Interpreter, value UInt16Value) {
	if v.UInt16ValueVisitor == nil {
		return
	}
	v.UInt16ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUInt32Value(interpreter *Interpreter, value UInt32Value) {
	if v.UInt32ValueVisitor == nil {
		return
	}
	v.UInt32ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUInt64Value(interpreter *Interpreter, value UInt64Value) {
	if v.UInt64ValueVisitor == nil {
		return
	}
	v.UInt64ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUInt128Value(interpreter *Interpreter, value UInt128Value) {
	if v.UInt128ValueVisitor == nil {
		return
	}
	v.UInt128ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUInt256Value(interpreter *Interpreter, value UInt256Value) {
	if v.UInt256ValueVisitor == nil {
		return
	}
	v.UInt256ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitWord8Value(interpreter *Interpreter, value Word8Value) {
	if v.Word8ValueVisitor == nil {
		return
	}
	v.Word8ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitWord16Value(interpreter *Interpreter, value Word16Value) {
	if v.Word16ValueVisitor == nil {
		return
	}
	v.Word16ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitWord32Value(interpreter *Interpreter, value Word32Value) {
	if v.Word32ValueVisitor == nil {
		return
	}
	v.Word32ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitWord64Value(interpreter *Interpreter, value Word64Value) {
	if v.Word64ValueVisitor == nil {
		return
	}
	v.Word64ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitFix64Value(interpreter *Interpreter, value Fix64Value) {
	if v.Fix64ValueVisitor == nil {
		return
	}
	v.Fix64ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitUFix64Value(interpreter *Interpreter, value UFix64Value) {
	if v.UFix64ValueVisitor == nil {
		return
	}
	v.UFix64ValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitCompositeValue(interpreter *Interpreter, value *CompositeValue) bool {
	if v.CompositeValueVisitor == nil {
		return true
	}
	return v.CompositeValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitDictionaryValue(interpreter *Interpreter, value *DictionaryValue) bool {
	if v.DictionaryValueVisitor == nil {
		return true
	}
	return v.DictionaryValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitNilValue(interpreter *Interpreter, value NilValue) {
	if v.NilValueVisitor == nil {
		return
	}
	v.NilValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitSomeValue(interpreter *Interpreter, value *SomeValue) bool {
	if v.SomeValueVisitor == nil {
		return true
	}
	return v.SomeValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitStorageReferenceValue(interpreter *Interpreter, value *StorageReferenceValue) {
	if v.StorageReferenceValueVisitor == nil {
		return
	}
	v.StorageReferenceValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitEphemeralReferenceValue(interpreter *Interpreter, value *EphemeralReferenceValue) {
	if v.EphemeralReferenceValueVisitor == nil {
		return
	}
	v.EphemeralReferenceValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitAddressValue(interpreter *Interpreter, value AddressValue) {
	if v.AddressValueVisitor == nil {
		return
	}
	v.AddressValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitAuthAccountValue(interpreter *Interpreter, value AuthAccountValue) {
	if v.AuthAccountValueVisitor == nil {
		return
	}
	v.AuthAccountValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitPublicAccountValue(interpreter *Interpreter, value PublicAccountValue) {
	if v.PublicAccountValueVisitor == nil {
		return
	}
	v.PublicAccountValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitPathValue(interpreter *Interpreter, value PathValue) {
	if v.PathValueVisitor == nil {
		return
	}
	v.PathValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitCapabilityValue(interpreter *Interpreter, value CapabilityValue) {
	if v.CapabilityValueVisitor == nil {
		return
	}
	v.CapabilityValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitLinkValue(interpreter *Interpreter, value LinkValue) {
	if v.LinkValueVisitor == nil {
		return
	}
	v.LinkValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitInterpretedFunctionValue(interpreter *Interpreter, value InterpretedFunctionValue) {
	if v.InterpretedFunctionValueVisitor == nil {
		return
	}
	v.InterpretedFunctionValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitHostFunctionValue(interpreter *Interpreter, value HostFunctionValue) {
	if v.HostFunctionValueVisitor == nil {
		return
	}
	v.HostFunctionValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitBoundFunctionValue(interpreter *Interpreter, value BoundFunctionValue) {
	if v.BoundFunctionValueVisitor == nil {
		return
	}
	v.BoundFunctionValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitAuthAccountContractsValue(interpreter *Interpreter, value AuthAccountContractsValue) {
	if v.AuthAccountContractsValueVisitor == nil {
		return
	}
	v.AuthAccountContractsValueVisitor(interpreter, value)
}

func (v EmptyVisitor) VisitDeployedContractValue(interpreter *Interpreter, value DeployedContractValue) {
	if v.DeployedContractValueVisitor == nil {
		return
	}
	v.DeployedContractValueVisitor(interpreter, value)
}