package AssetName

import (
	"encoding/hex"
	"errors"

	"github.com/Salvionied/cbor/v2"
)

type AssetName struct {
	Value string   `plutusType:"HexString"`
}

/*
internal use only
*/
func NewAssetNameFromHexString(value string) *AssetName {
	_, err := hex.DecodeString(value)

	if err != nil || len(value) > 64 {
		return nil
	}

	return &AssetName{Value: value}
}

func NewAssetNameFromString(value string) AssetName {
	v := hex.EncodeToString([]byte(value))
	return AssetName{Value: v}
}

func (an AssetName) String() string {
	decoded, _ := hex.DecodeString(an.Value)
	return string(decoded)
}

func (an AssetName) HexString() string {
	return an.Value
}

func (an *AssetName) MarshalCBOR() ([]byte, error) {
	if an.Value == "[]" || an.Value == "" {
		return cbor.Marshal(make([]byte, 0))
	}

	if len(an.Value) > 64 {
		return nil, errors.New("invalid asset name length")
	}

	byteSlice, _ := hex.DecodeString(an.Value)

	return cbor.Marshal(byteSlice)
}

func (an *AssetName) UnmarshalCBOR(value []byte) error {
	var res []byte
	err := cbor.Unmarshal(value, &res)
	if err != nil {
		return err
	}

	if len(res) > 32 {
		return errors.New("invalid asset name length")
	}

	an.Value = hex.EncodeToString(res)

	return nil
}
