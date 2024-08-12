package Fingerprint

import (
	"encoding/hex"

	"github.com/Salvionied/apollo/crypto/bech32"
	"github.com/Salvionied/apollo/plutusencoder"
	"github.com/Salvionied/apollo/serialization/AssetName"
	"github.com/Salvionied/apollo/serialization/PlutusData"
	"github.com/Salvionied/apollo/serialization/Policy"
	"golang.org/x/crypto/blake2b"
)

type Fingerprint struct {
	_         struct{} `plutusType:"IndefList" plutusConstr:"0"`
	PolicyId  Policy.PolicyId
	AssetName AssetName.AssetName
}

func New(policyId Policy.PolicyId, assetName AssetName.AssetName) *Fingerprint {
	return &Fingerprint{
		PolicyId:  policyId,
		AssetName: assetName,
	}
}


func (f *Fingerprint) String() string {
	bs, _ := hex.DecodeString(f.PolicyId.Value + f.AssetName.HexString())
	hasher, _ := blake2b.New(20, nil)
	hasher.Write(bs)
	hashBytes := hasher.Sum(nil)

	words, _ := bech32.ConvertBits(hashBytes, 8, 5, false)
	result, _ := bech32.Encode("asset", words)
	return result
}

func (f Fingerprint) ToPlutusData() (PlutusData.PlutusData, error) {
	result, err := plutusencoder.MarshalPlutus(f)
	if err != nil {
		return PlutusData.PlutusData{}, err
	}
	return *result, nil
}
