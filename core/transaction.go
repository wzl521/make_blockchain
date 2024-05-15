package core

import (
	"fmt"
	"make-blockchain/crypto"
)

// Transaction 表示区块链中的交易结构体，包含交易数据、公钥和签名信息。
type Transaction struct {
	Data      []byte            // 交易数据
	PublicKey crypto.PublicKey  // 交易的公钥
	Signature *crypto.Signature // 交易的签名
}

// Sign 使用给定的私钥对交易进行签名。
func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	// 使用私钥对交易数据进行签名
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	// 设置交易的公钥为私钥对应的公钥
	tx.PublicKey = privKey.PublicKey()
	// 设置交易的签名
	tx.Signature = sig
	return nil
}

// Verify 验证交易的签名是否有效。
func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}
	// 验证交易的签名是否有效
	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}
	return nil
}
