package types

import (
	"encoding/hex"
	"fmt"
)

// Address 表示一个长度为 20 的字节数组类型，用于表示地址。
type Address [20]uint8

// ToSlice 将地址转换为字节切片。
func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

// AddressFromBytes 根据给定的字节切片创建一个地址。
// 字节切片的长度必须为 20，否则会触发 panic。
func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("given bytes with length %d should be 20", len(b))
		panic(msg)
	}
	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}
	return Address(value)
}

// String 返回地址的十六进制表示形式。
func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}
