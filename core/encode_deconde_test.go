package core

import (
	"bytes"
	"time"

	"testing"

	"github.com/hosseinal/BlockChain/types"

	"github.com/stretchr/testify/assert"
)

func TestHeaderEncodeDecode(t *testing.T) {
	originalHeader := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Heght:     100,
		Nonce:     12345,
	}

	buf := bytes.Buffer{}
	err := originalHeader.EncodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	decodedHeader := &Header{}
	err = decodedHeader.DecodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, originalHeader.Version, decodedHeader.Version)
	assert.Equal(t, originalHeader.PrevBlock, decodedHeader.PrevBlock)
	assert.Equal(t, originalHeader.Timestamp, decodedHeader.Timestamp)
	assert.Equal(t, originalHeader.Heght, decodedHeader.Heght)
	assert.Equal(t, originalHeader.Nonce, decodedHeader.Nonce)
}
