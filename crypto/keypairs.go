package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/hosseinal/BlockChain/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GetPrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		panic(err)
	}

	return PrivateKey{key: key}
}

func (pk PrivateKey) PublicKey() PublicKey {
	return PublicKey{key: &pk.key.PublicKey}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pk PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(pk.key, pk.key.X, pk.key.Y)
}

func (pk PublicKey) Address() types.Address {

	h := sha256.Sum256(pk.ToSlice())

	return types.AddressFromBytes(h[:20])
}

type Signature struct {
}
