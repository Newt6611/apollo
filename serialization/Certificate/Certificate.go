package Certificate

import "github.com/Newt6611/apollo/serialization"

type StakeCredential struct {
	_          struct{} `cbor:"toarray"`
	_CODE      int      `cbor:",omitempty"`
	Credential serialization.ConstrainedBytes
}

// TODO
type Certificate struct {
	_               struct{} `cbor:"toarray"`
	_CODE           int
	StakeCredential *StakeCredential
}

type Certificates []*Certificate
