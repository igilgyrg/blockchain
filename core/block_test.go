package core

import (
	"bytes"
	"math/rand"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/igilgyrg/blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_EncodeBinary_DecodeBinary(t *testing.T) {
	h := RandomHeader()

	buf := &bytes.Buffer{}
	err := h.EncodeBinary(buf)
	assert.NoError(t, err)

	hDecode := &Header{}
	err = hDecode.DecodeBinary(buf)
	assert.NoError(t, err)

	assert.Equal(t, h, hDecode)
}

func TestBlock_EncodeBinary_DecodeBinary(t *testing.T) {
	b := &Block{
		Header:       *RandomHeader(),
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	err := b.EncodeBinary(buf)
	assert.NoError(t, err)

	bDecode := &Block{}
	err = bDecode.DecodeBinary(buf)
	assert.NoError(t, err)

	assert.Equal(t, b, bDecode)
}

func TestBlock_Hash(t *testing.T) {
	h := RandomHeader()
	b1 := &Block{
		Header:       *h,
		Transactions: []Transaction{},
	}

	b2 := &Block{
		Header:       *h,
		Transactions: []Transaction{},
	}

	b1Hash := b1.Hash()
	assert.False(t, b1Hash.IsZero())
	assert.Equal(t, b1Hash, b1.Hash())
	assert.Equal(t, b1Hash, b2.Hash())
}
func RandomHeader() *Header {
	return &Header{
		Version:   gofakeit.Uint64(),
		PrevBlock: RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    gofakeit.Uint32(),
		Nonce:     gofakeit.Uint64(),
	}
}

func RandomHash() types.Hash {
	hashBytes := make([]byte, 32)
	rand.Read(hashBytes)
	return types.HashFromBytes(hashBytes)
}
