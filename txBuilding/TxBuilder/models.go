package TxBuilder

import (
	"github.com/Newt6611/apollo/serialization/PlutusData"
	"github.com/Newt6611/apollo/serialization/Redeemer"
)

type MintingScriptToRedeemer struct {
	Script      PlutusData.ScriptHashable
	Redeemer    Redeemer.Redeemer
	HasRedeemer bool
}
