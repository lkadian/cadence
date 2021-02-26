/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sema

import "github.com/onflow/cadence/runtime/errors"

//go:generate go run golang.org/x/tools/cmd/stringer -type=SignatureAlgorithm
//go:generate go run golang.org/x/tools/cmd/stringer -type=HashAlgorithm

var SignatureAlgorithms = []BuiltinEnumCase{
	SignatureAlgorithmECDSA_P256,
	SignatureAlgorithmECDSA_Secp256k1,
	SignatureAlgorithmBLSBLS12381,
}

var HashAlgorithms = []BuiltinEnumCase{
	HashAlgorithmSHA2_256,
	HashAlgorithmSHA2_384,
	HashAlgorithmSHA3_256,
	HashAlgorithmSHA3_384,
	HashAlgorithmKMAC128,
}

var SignatureAlgorithmType = newEnumType(SignatureAlgorithmTypeName, &UInt8Type{})

type SignatureAlgorithm int

const (
	// Supported signing algorithms
	SignatureAlgorithmECDSA_P256 SignatureAlgorithm = iota
	SignatureAlgorithmECDSA_Secp256k1
	SignatureAlgorithmBLSBLS12381
)

// Name returns the string representation of this signing algorithm.
func (algo SignatureAlgorithm) Name() string {
	switch algo {
	case SignatureAlgorithmECDSA_P256:
		return "ECDSA_P256"
	case SignatureAlgorithmECDSA_Secp256k1:
		return "ECDSA_Secp256k1"
	case SignatureAlgorithmBLSBLS12381:
		return "BLSBLS12381"
	}

	panic(errors.NewUnreachableError())
}

func (algo SignatureAlgorithm) RawValue() int {
	return int(algo)
}

func (algo SignatureAlgorithm) DocString() string {
	switch algo {
	case SignatureAlgorithmECDSA_P256:
		return SignatureAlgorithmDocStringECDSA_P256
	case SignatureAlgorithmECDSA_Secp256k1:
		return SignatureAlgorithmDocStringECDSA_Secp256k1
	case SignatureAlgorithmBLSBLS12381:
		return SignatureAlgorithmDocStringBLSBLS12381
	}

	panic(errors.NewUnreachableError())
}

const SignatureAlgorithmTypeName = "SignatureAlgorithm"

const SignatureAlgorithmDocStringECDSA_P256 = `
SignatureAlgorithmECDSA_P256 is Elliptic Curve Digital Signature Algorithm (ECDSA) on the NIST P-256 curve
`
const SignatureAlgorithmDocStringECDSA_Secp256k1 = `
SignatureAlgorithmECDSA_Secp256k1 is Elliptic Curve Digital Signature Algorithm (ECDSA) on the secp256k1 curve
`
const SignatureAlgorithmDocStringBLSBLS12381 = `
SignatureAlgorithmDocStringBLSBLS12381 is BLS Signature algorithm on BLS 12-381 curve
`

const HashAlgorithmTypeName = "HashAlgorithm"

const HashAlgorithmDocStringSHA2_256 = `
HashAlgorithmSHA2_256 is Secure Hashing Algorithm 2 (SHA-2) with a 256-bit digest
`
const HashAlgorithmDocStringSHA2_384 = `
HashAlgorithmSHA2_384 is Secure Hashing Algorithm 2 (SHA-2) with a 384-bit digest
`

const HashAlgorithmDocStringSHA3_256 = `
HashAlgorithmSHA3_256 is Secure Hashing Algorithm 3 (SHA-3) with a 256-bit digest
`

const HashAlgorithmDocStringSHA3_384 = `
HashAlgorithmSHA3_384 is Secure Hashing Algorithm 3 (SHA-3) with a 384-bit digest
`

const HashAlgorithmDocStringKMAC128 = `
HashAlgorithmKMAC128 is KECCAK Message Authentication Code with a 128-bit digest
`

var HashAlgorithmType = newEnumType(HashAlgorithmTypeName, &UInt8Type{})

type HashAlgorithm int

const (
	// Supported hashing algorithms
	HashAlgorithmSHA2_256 HashAlgorithm = iota
	HashAlgorithmSHA2_384
	HashAlgorithmSHA3_256
	HashAlgorithmSHA3_384
	HashAlgorithmKMAC128
)

func (algo HashAlgorithm) Name() string {
	switch algo {
	case HashAlgorithmSHA2_256:
		return "SHA2_256"
	case HashAlgorithmSHA2_384:
		return "SHA2_384"
	case HashAlgorithmSHA3_256:
		return "SHA3_256"
	case HashAlgorithmSHA3_384:
		return "SHA3_384"
	case HashAlgorithmKMAC128:
		return "KMAC128"
	}

	panic(errors.NewUnreachableError())
}

func (algo HashAlgorithm) RawValue() int {
	return int(algo)
}

func (algo HashAlgorithm) DocString() string {
	switch algo {
	case HashAlgorithmSHA2_256:
		return HashAlgorithmDocStringSHA2_256
	case HashAlgorithmSHA2_384:
		return HashAlgorithmDocStringSHA2_384
	case HashAlgorithmSHA3_256:
		return HashAlgorithmDocStringSHA3_256
	case HashAlgorithmSHA3_384:
		return HashAlgorithmDocStringSHA3_384
	case HashAlgorithmKMAC128:
		return HashAlgorithmDocStringKMAC128
	}

	panic(errors.NewUnreachableError())
}

func newEnumType(identifier string, rawType Type) *BuiltinCompositeType {
	accountKeyType := &BuiltinCompositeType{
		Identifier:           identifier,
		EnumRawType:          rawType,
		IsInvalid:            false,
		IsResource:           false,
		Storable:             true,
		Equatable:            true,
		ExternallyReturnable: true,
	}

	// Members of the enum type are *not* the enum cases!
	// Each individual enum case is an instance of the enum type,
	// so only has a single member, the raw value field
	var members = []*Member{
		NewPublicEnumCaseMember(
			rawType,
			EnumRawValueFieldName,
			enumRawValueFieldDocString,
		),
	}

	accountKeyType.Members = GetMembersAsMap(members)
	return accountKeyType
}
