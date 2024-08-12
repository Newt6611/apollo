package TransactionWitnessSet_test

import (
	"encoding/hex"
	"testing"

	"github.com/Newt6611/apollo/serialization/PlutusData"
	"github.com/Newt6611/apollo/serialization/TransactionWitnessSet"
	"github.com/Salvionied/cbor/v2"
)

func TestMarshalAndUnmarshalNoScripts(t *testing.T) {
	tws := TransactionWitnessSet.TransactionWitnessSet{}
	twsBytes, err := cbor.Marshal(tws)
	if err != nil {
		t.Errorf("Error marshaling TransactionWitnessSet: %v", err)
	}
	if hex.EncodeToString(twsBytes) != "a1049fff" {
		t.Error("TransactionWitnessSet marshaled incorrectly", hex.EncodeToString(twsBytes))
	}
}

var pd = PlutusData.PlutusData{
	PlutusDataType: PlutusData.PlutusBytes,
	TagNr:          0,
	Value:          []byte{0x01, 0x02, 0x03},
}

func TestMarshalBasicPlutus(t *testing.T) {
	tws := TransactionWitnessSet.TransactionWitnessSet{
		PlutusData: PlutusData.PlutusIndefArray{pd},
	}
	twsBytes, err := cbor.Marshal(tws)
	if err != nil {
		t.Errorf("Error marshaling TransactionWitnessSet: %v", err)
	}
	if hex.EncodeToString(twsBytes) != "a1049f43010203ff" {
		t.Error("TransactionWitnessSet marshaled incorrectly", hex.EncodeToString(twsBytes))
	}
}
