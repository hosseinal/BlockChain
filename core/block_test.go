package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/hosseinal/BlockChain/types"
	"github.com/stretchr/testify/assert"
)

func TestBlockEncodeDecode(t *testing.T) {
	originalBlock := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Heght:     100,
			Nonce:     12345,
		},
		Transactions: nil,
	}

	buf := bytes.Buffer{}
	err := originalBlock.EncodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	decodedBlock := &Block{}
	err = decodedBlock.DecodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, originalBlock.Header.Version, decodedBlock.Header.Version)
	assert.Equal(t, originalBlock.Header.PrevBlock, decodedBlock.Header.PrevBlock)
	assert.Equal(t, originalBlock.Header.Timestamp, decodedBlock.Header.Timestamp)
	assert.Equal(t, originalBlock.Header.Heght, decodedBlock.Header.Heght)
	assert.Equal(t, originalBlock.Header.Nonce, decodedBlock.Header.Nonce)
	assert.Equal(t, len(originalBlock.Transactions), len(decodedBlock.Transactions))
}
