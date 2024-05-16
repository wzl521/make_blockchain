package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// newBlockchainWithGenesis函数创建一个带有创世区块的区块链实例，用于测试
// 它接受一个testing.T指针，并返回一个指向Blockchain的指针
func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	// 创建一个带有随机创世区块的区块链实例，并检查是否没有错误发生
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	// 返回创建的区块链实例
	return bc
}

// TestAddBlock函数测试向区块链添加区块的功能
func TestAddBlock(t *testing.T) {
	// 创建一个带有创世区块的区块链实例
	bc := newBlockchainWithGenesis(t)

	// 定义要添加的区块数量
	lenBlocks := 1000
	// 循环添加1000个区块到区块链中
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))
		// 检查添加区块时是否没有错误发生
		assert.Nil(t, bc.AddBlock(block))
	}

	// 检查区块链的高度是否等于添加的区块数量
	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	// 检查headers切片的长度是否等于添加的区块数量加1（因为包括创世区块）
	assert.Equal(t, len(bc.headers), lenBlocks+1)
	// 尝试添加一个无效的区块，应该返回非nil的错误
	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}

// TestNewBlockchain函数测试创建新的区块链实例的功能
func TestNewBlockchain(t *testing.T) {
	// 创建一个带有创世区块的区块链实例
	bc := newBlockchainWithGenesis(t)
	// 检查区块链的验证器是否不为nil
	assert.NotNil(t, bc.validator)
	// 检查区块链的高度是否为0（因为只有创世区块）
	assert.Equal(t, bc.Height(), uint32(0))
}

// TestHasBlock函数测试检查区块链是否包含特定高度区块的功能
func TestHasBlock(t *testing.T) {
	// 创建一个带有创世区块的区块链实例
	bc := newBlockchainWithGenesis(t)
	// 检查区块链是否包含高度为0的区块（即创世区块）
	assert.True(t, bc.HasBlock(0))
}
