package core

import "io"

// Encoder[T] 是一个泛型接口，用于将类型 T 的对象编码为字节流并写入到 io.Writer。
type Encoder[T any] interface {
	Encode(io.Writer, T) error
}

// Decoder[T] 是一个泛型接口，用于从 io.Reader 中读取字节流并解码为类型 T 的对象。
type Decoder[T any] interface {
	Decode(io.Reader, T) error
}
