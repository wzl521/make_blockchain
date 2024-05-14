package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalTransport_Connect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, trb.peers[trb.addr], tra)
}
func TestLocalTransport_SendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.consumeCh
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.addr)

}
