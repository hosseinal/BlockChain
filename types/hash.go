package types

type Hash [32]uint8

func HashFromBytes(data []byte) Hash {
	if len(data) != 32 {
		panic("invalid hash length")
	}

	var h Hash
	copy(h[:], data)
	return h
}
