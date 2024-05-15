package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert" // 导入断言库，用于测试断言
	"testing"
)

// TestKeyPairSignVerifySuccess 测试密钥对签名和验证成功的情况。
func TestKeyPairSignVerifySuccess(t *testing.T) {
	// 生成私钥和对应的公钥
	privateKey := GeneratePrivateKey()
	pubKey := privateKey.PublicKey()

	// 准备要签名的消息
	msg := []byte("hello world")

	// 使用私钥对消息进行签名
	sig, err := privateKey.Sign(msg)
	assert.Nil(t, err) // 使用断言库断言签名操作没有错误

	fmt.Println(sig) // 打印签名结果

	// 使用公钥验证签名是否有效
	assert.True(t, sig.Verify(pubKey, msg)) // 使用断言库断言签名验证成功
}

// TestKeyPairSignVerifyFail 测试密钥对签名和验证失败的情况。
func TestKeyPairSignVerifyFail(t *testing.T) {
	// 生成第一个私钥和对应的公钥
	privateKey := GeneratePrivateKey()
	pubKey := privateKey.PublicKey()

	// 准备要签名的消息
	msg := []byte("hello world")

	// 使用第一个私钥对消息进行签名
	sig, err := privateKey.Sign(msg)
	assert.Nil(t, err) // 使用断言库断言签名操作没有错误

	// 生成第二个私钥和对应的公钥
	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	// 使用第二个公钥验证签名，预期会失败
	assert.False(t, sig.Verify(otherPublicKey, msg)) // 使用断言库断言签名验证失败

	// 使用第一个公钥验证错误的消息，预期会失败
	assert.False(t, sig.Verify(pubKey, []byte("xxxxxxx"))) // 使用断言库断言签名验证失败
}
