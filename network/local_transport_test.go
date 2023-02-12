package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	err := tra.Connect(trb)
	assert.NoError(t, err)

	err = trb.Connect(tra)
	assert.NoError(t, err)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	err := tra.Connect(trb)
	assert.NoError(t, err)

	err = trb.Connect(tra)
	assert.NoError(t, err)

	msgTra := []byte("some message")
	err = tra.SendMessage(trb.Addr(), msgTra)
	assert.NoError(t, err)

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msgTra)

	msgTrb := []byte("another message")
	err = trb.SendMessage(tra.Addr(), msgTrb)
	assert.NoError(t, err)

	rpc = <-tra.Consume()
	assert.Equal(t, rpc.Payload, msgTrb)
}
