package VerificationKeyWitness

import "github.com/Newt6611/apollo/serialization/Key"

type VerificationKeyWitness struct {
	_         struct{} `cbor:",toarray"`
	Vkey      Key.VerificationKey
	Signature []uint8
}
