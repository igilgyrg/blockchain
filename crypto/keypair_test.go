package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKeyValid(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubKey := privkey.PublicKey()

	msg := []byte("hello world once again")
	sign, err := privkey.Sign(msg)
	assert.NoError(t, err)
	assert.NotNil(t, sign)

	verify := sign.Verify(pubKey, msg)
	assert.True(t, verify)
}

func TestGeneratePrivateKeyInvalid(t *testing.T) {
	privkey := GeneratePrivateKey()

	msg := []byte("hello world once again")
	sign, err := privkey.Sign(msg)
	assert.NoError(t, err)
	assert.NotNil(t, sign)

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKet := otherPrivateKey.PublicKey()

	verify := sign.Verify(otherPublicKet, msg)
	assert.False(t, verify)
}
