package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"make-blockchain/crypto"
	"make-blockchain/types"
	"testing"
	"time"
)

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height)
	assert.Nil(t, b.Sign(privKey))

	return b
}
func randomBlock(height uint32) *Block {
	header := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Height:    height,
		Timestamp: time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("hello"),
	}
	return NewBlock(header, []Transaction{tx})
}
func TestHashBlock(t *testing.T) {
	b := randomBlock(100)
	fmt.Println(b.Hash(BlockHasher{}))
}
func TestBlock_Sign(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestBlock_Verify(t *testing.T) {
	// 生成一个新的私钥作为区块签名的私钥
	privKey := crypto.GeneratePrivateKey()

	// 随机生成一个区块
	b := randomBlock(0)

	// 对区块进行签名，使用生成的私钥进行签名
	assert.Nil(t, b.Sign(privKey))

	// 验证区块的签名是否有效
	assert.Nil(t, b.Verify())

	// 创建一个空的私钥，作为区块的验证人公钥
	otherPrivKey := crypto.GeneratePrivateKey()

	// 将区块的验证人设为这个空的私钥对应的公钥
	b.Validator = otherPrivKey.PublicKey()

	// 验证区块的签名是否有效，预期会失败，因为验证人公钥为空
	assert.NotNil(t, b.Verify())

	// 修改区块的高度为 10
	b.Height = 100

	// 验证区块的签名是否有效，预期会失败，因为区块高度发生了变化
	assert.NotNil(t, b.Verify())
}
