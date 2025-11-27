package types

import "encoding/hex"

type Address [20]byte

func (a Address) IsZero() bool {
	for _, b := range a {
		if b != 0 {
			return false
		}
	}
	return true
}

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	copy(b, a[:])
	return b
}

func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

func AddressFromBytes(data []byte) Address {
	if len(data) != 20 {
		panic("invalid address length")
	}

	var a Address
	copy(a[:], data)
	return a
}
