package core

import (
	"crypto/sha256"
	"make-blockchain/types"
)

// Hasher[T] 是一个泛型接口，用于计算类型 T 的对象的哈希值。
type Hasher[T any] interface {
	Hash(T) types.Hash
}

// BlockHasher 是一个实现了 Hasher[Block] 接口的具体类型，用于计算 Block 类型对象的哈希值。
type BlockHasher struct {
}

// Hash 实现了 Hasher[Block] 接口，用于计算 Block 类型对象的哈希值。
func (BlockHasher) Hash(b *Block) types.Hash {
	h := sha256.Sum256(b.HeaderData()) // 计算区块头部数据的 SHA-256 哈希值
	return types.Hash(h)               // 将哈希值转换为 types.Hash 类型并返回
}
