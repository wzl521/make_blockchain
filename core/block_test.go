package core

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"make-blockchain/types"
	"testing"
	"time"
)

func TestHeader_Encode_Decode_Binary(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    30,
		Nonce:     934934044,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode_Binary(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    30,
			Nonce:     934934044,
		},
		Transactions: nil,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
	fmt.Printf("%+v", bDecode)

}

func TestBlock_Hash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    30,
			Nonce:     934934044,
		},
		Transactions: nil,
	}
	h := b.Hash()
	fmt.Println(h)
	assert.False(t, h.IsZero())
}
