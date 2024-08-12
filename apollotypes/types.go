package apollotypes

import (
	"github.com/Newt6611/apollo/serialization"
	serAddress "github.com/Newt6611/apollo/serialization/Address"
	"github.com/Newt6611/apollo/serialization/Key"
	"github.com/Newt6611/apollo/serialization/Transaction"
	"github.com/Newt6611/apollo/serialization/TransactionWitnessSet"
	"github.com/Newt6611/apollo/serialization/VerificationKeyWitness"
	"github.com/Newt6611/apollo/txBuilding/Backend/Base"
)

type Wallet interface {
	GetAddress() *serAddress.Address
	SignTx(tx Transaction.Transaction) TransactionWitnessSet.TransactionWitnessSet
	PkeyHash() serialization.PubKeyHash
	//SignMessage(address serAddress.Address, message []uint8) []uint8
}

type ExternalWallet struct {
	Address serAddress.Address
}

/**
	GetAddress returns the address associated with an external wallet.

	Returns:
		*serAddress.Address: A pointer to the address of the external wallet.
*/
func (ew *ExternalWallet) GetAddress() *serAddress.Address {
	return &ew.Address
}

/**
	SignTx signs a transaction using an external wallet.

	Params:
		tx (Transaction.Transaction): The transaction to be signed.

	Returns:
		TransactionWitnessSet.TransactionWitnessSet: The withness set associated with the signed transaction.
*/
func (ew *ExternalWallet) SignTx(tx Transaction.Transaction) TransactionWitnessSet.TransactionWitnessSet {
	return tx.TransactionWitnessSet
}

/**
	PkeyHash returns the public key hash assoicated with an external wallet.
	It computes and returns the public key hash based on the PaymentPart 
	of the wallet's address.

	Returns:
		serialization.PubKeyHash: The public key hash of the external wallet.
*/
func (ew *ExternalWallet) PkeyHash() serialization.PubKeyHash {
	res := serialization.PubKeyHash(ew.Address.PaymentPart)
	return res
}

type GenericWallet struct {
	SigningKey           Key.SigningKey
	VerificationKey      Key.VerificationKey
	Address              serAddress.Address
	StakeSigningKey      Key.StakeSigningKey
	StakeVerificationKey Key.StakeVerificationKey
}

/**
	PkeyHash calculates and returns the public key hash associated with a generic wallet.
	It computes the public key hash by calling the Hash() method on the wallet's VerificationKey.
	Then it returns as a serialization.PubKeyHas type.
	
	Returns:
   		serialization.PubKeyHash: The public key hash of the generic wallet.
*/
func (gw *GenericWallet) PkeyHash() serialization.PubKeyHash {
	res, _ := gw.VerificationKey.Hash()
	return res
}

/**
	GetAddress returns the address associated with a generic wallet.

	Returns:
		*serAddress.Address: A pointer to the address of a generic wallet.
*/
func (gw *GenericWallet) GetAddress() *serAddress.Address {
	return &gw.Address
}

/**
	SignTx signs a transaction using a generic wallet and returns the updated TransactionWitnessSet.
	It takes a transaction of type Transaction.Transaction and signs it using the wallet's SigningKey.
	Then it appends the corresponding VerificationKeyWitness to the TransactionWitnessSet and returns
	the updated witness set.

	Parameters:
	   	wallet (*GenericWallet): A pointer to a generic wallet.
		tx (Transaction.Transaction): The transaction to be signed.

	Returns:
   		TransactionWitnessSet.TransactionWitnessSet: The updated TransactionWitnessSet after signing the transaction.
*/
func (wallet *GenericWallet) SignTx(tx Transaction.Transaction) TransactionWitnessSet.TransactionWitnessSet {
	witness_set := tx.TransactionWitnessSet
	txHash, _ := tx.TransactionBody.Hash()
	signature, _ := wallet.SigningKey.Sign(txHash)
	witness_set.VkeyWitnesses = append(witness_set.VkeyWitnesses, VerificationKeyWitness.VerificationKeyWitness{Vkey: wallet.VerificationKey, Signature: signature})
	return witness_set
}

type Backend Base.ChainContext

type Address serAddress.Address
