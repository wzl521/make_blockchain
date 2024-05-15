package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"make-blockchain/types"
	"math/big"
)

// PrivateKey 表示 ECDSA 加密算法的私钥结构。
type PrivateKey struct {
	key *ecdsa.PrivateKey // ECDSA 私钥
}

// GeneratePrivateKey 生成一个新的 ECDSA 私钥。
func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // 使用 P-256 曲线生成私钥
	if err != nil {
		panic(err)
	}
	return PrivateKey{
		key: key,
	}
}

// PublicKey 表示 ECDSA 加密算法的公钥结构。
type PublicKey struct {
	key *ecdsa.PublicKey // ECDSA 公钥
}

// PublicKey 方法返回与私钥关联的公钥。
func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey, // 获取私钥对应的公钥
	}
}

// ToSlice 将公钥转换为压缩格式的字节切片。
func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y) // 将公钥按压缩格式转换为字节切片
}

// Address 计算公钥的地址，使用 SHA-256 哈希算法。
// 地址是哈希值的后 20 字节。
func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())              // 计算公钥的 SHA-256 哈希值
	return types.AddressFromBytes(h[len(h)-20:]) // 将哈希值的后 20 字节转换为地址类型
}

// Signature 表示 ECDSA 签名结构体。
type Signature struct {
	r, s *big.Int // ECDSA 签名的 r 和 s 值
}

// Sign 使用私钥对指定数据进行签名，并返回签名结果。
func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data) // 使用私钥对数据进行签名
	if err != nil {
		return nil, err
	}
	return &Signature{r, s}, nil // 返回签名结果
}

// Verify 验证给定公钥对指定数据的签名是否有效。
func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.key, data, sig.r, sig.s) // 使用公钥验证签名的有效性
}
