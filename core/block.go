package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"make-blockchain/crypto"
	"make-blockchain/types"
)

// Header 定义区块的头部信息，包括版本号、上一个区块的哈希、时间戳、高度和随机数。
type Header struct {
	Version   uint32     // 区块版本号
	PrevBlock types.Hash // 上一个区块的哈希值
	Timestamp int64      // 区块的时间戳
	Height    uint32     // 区块的高度
	Nonce     uint64     // 区块的随机数
}

// Block 定义区块结构，包括区块头部信息、交易列表、验证人公钥和区块签名。
type Block struct {
	*Header                        // 指向区块的头部信息
	Transactions []Transaction     // 区块包含的交易列表
	Validator    crypto.PublicKey  // 区块的验证人公钥
	Signature    *crypto.Signature // 区块的签名

	// 缓存的区块哈希值
	hash types.Hash
}

// NewBlock 创建一个新的区块实例。
func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

// Sign 对区块进行签名，使用给定的私钥。
func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData()) // 使用区块头部数据进行签名
	if err != nil {
		return err
	}
	b.Validator = privKey.PublicKey() // 设置区块的验证人公钥为私钥对应的公钥
	b.Signature = sig                 // 设置区块的签名
	return nil
}

// Verify 验证区块的签名是否有效。
func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}
	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("block has invalid signature")
	}
	return nil
}

// Decode 从输入流中解码区块数据。
func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

// Encode 将区块数据编码到输出流中。
func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

// Hash 计算区块的哈希值，使用给定的哈希函数。
func (b *Block) Hash(harsher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = harsher.Hash(b)
	}
	return b.hash
}

// HeaderData 将区块头部信息编码为字节数据。
func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)
	return buf.Bytes()
}
