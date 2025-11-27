package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPair(t *testing.T) {
	privKey := GetPrivateKey()
	pubKey := privKey.PublicKey()

	pubKeySlice := pubKey.ToSlice()
	if len(pubKeySlice) == 0 {
		t.Errorf("Public key slice is empty")
	}

	address := pubKey.Address()
	if address.IsZero() {
		t.Errorf("Address derived from public key is zero")
	}

	fmt.Printf("Public Key: %x\n", pubKeySlice)
	fmt.Printf("Address: %x\n", address)

	signature, err := privKey.Sign([]byte("test data"))
	if err != nil {
		t.Errorf("Failed to sign data: %v", err)
	}

	if signature.r == nil || signature.s == nil {
		t.Errorf("Signature components are nil")
	}

	assert.True(t, signature.Verify(pubKey, []byte("test data")))
}

func TestKeyPairFailure(t *testing.T) {
	privKey := GetPrivateKey()
	pubKey := privKey.PublicKey()

	otherPrivKey := GetPrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	signature, err := privKey.Sign([]byte("test data"))
	if err != nil {
		t.Errorf("Failed to sign data: %v", err)
	}

	assert.False(t, signature.Verify(otherPubKey, []byte("test data")))
	assert.False(t, signature.Verify(pubKey, []byte("different data")))
	fmt.Println("All failure cases passed as expected")
}
