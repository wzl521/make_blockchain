package core // 定义当前包为core

// Blockchain结构体代表一个区块链实例
type Blockchain struct {
	// store字段是一个Storage接口，用于存储区块链数据
	store Storage
	// headers字段是一个Header指针切片，用于存储区块链的区块头
	headers []*Header
	// validator字段是一个Validator接口，用于验证区块的有效性
	validator Validator
}

// NewBlockchain函数创建一个新的区块链实例，并添加创世区块
// 它返回一个指向Blockchain的指针和一个可能的错误
func NewBlockchain(genesis *Block) (*Blockchain, error) {
	// 初始化一个新的Blockchain实例，初始headers为空切片，store为内存存储
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemorystore(),
	}
	// 设置区块链的验证器为一个新的BlockValidator实例，传入当前区块链实例
	bc.validator = NewBlockValidator(bc)
	// 尝试将创世区块添加到区块链中，不进行验证
	err := bc.addBlockWithoutValidation(genesis)

	// 返回区块链实例和可能的错误
	return bc, err
}

// SetValidator方法用于设置区块链的验证器
func (bc *Blockchain) SetValidator(v Validator) {
	// 将区块链的验证器字段设置为传入的验证器实例
	bc.validator = v
}

// AddBlock方法用于向区块链添加一个新的区块
// 在添加之前，它会使用区块链的验证器来验证区块的有效性
func (bc *Blockchain) AddBlock(b *Block) error {
	// 验证区块，如果验证失败，返回错误
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	// 如果验证成功，将区块添加到区块链中，不进行验证
	return bc.addBlockWithoutValidation(b)
}

// HasBlock方法检查区块链是否包含指定高度的区块
func (bc *Blockchain) HasBlock(height uint32) bool {
	// 如果指定的高度小于或等于区块链的高度，则返回true
	return height <= bc.Height()
}

// Height方法返回区块链的高度，即最后一个区块的索引
func (bc *Blockchain) Height() uint32 {
	// 区块链的高度等于headers切片的长度减1
	return uint32(len(bc.headers) - 1)
}

// addBlockWithoutValidation方法用于将区块添加到区块链中，但不进行验证
func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	// 将区块的Header添加到headers切片中
	bc.headers = append(bc.headers, b.Header)
	// 将区块存储到store中
	return bc.store.Put(b)
}
