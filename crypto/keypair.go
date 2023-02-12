package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/igilgyrg/blockchain/types"
)

type (
	PrivateKey struct {
		key *ecdsa.PrivateKey
	}

	PublicKey struct {
		key *ecdsa.PublicKey
	}

	Signature struct {
		r, s *big.Int
	}
)

func (p PrivateKey) PublicKey() PublicKey {
	return PublicKey{key: &p.key.PublicKey}
}

func (p PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, p.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{r, s}, nil
}

func (p PublicKey) ToSlice() []uint8 {
	return elliptic.MarshalCompressed(p.key, p.key.X, p.key.Y)
}

func (p PublicKey) Address() types.Address {
	h := sha256.Sum256(p.ToSlice())
	return types.AddressFromBytes(h[len(h)-20:])
}

func (s *Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.key, data, s.r, s.s)
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic("error generate private key")
	}

	return PrivateKey{key: key}
}
