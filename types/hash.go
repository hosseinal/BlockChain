package types

import "encoding/hex"

type Hash [32]uint8

func (h Hash) IsZero() bool {
	for _, b := range h {
		if b != 0 {
			return false
		}
	}
	return true
}

func (h Hash) ToSlice() []byte {
	b := make([]byte, 32)
	copy(b, h[:])
	return b
}

func HashFromBytes(data []byte) Hash {
	if len(data) != 32 {
		panic("invalid hash length")
	}

	var h Hash
	copy(h[:], data)
	return h
}

func ToHash(data []byte) Hash {
	return HashFromBytes(data)
}

func (h Hash) ToString() string {
	return hex.EncodeToString(h.ToSlice())
}

func RandomByte(size int) []byte {
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = byte(i % 256)
	}
	return b
}

func RandomHash() Hash {
	return HashFromBytes(RandomByte(32))
}
