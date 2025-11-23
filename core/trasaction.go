package core

import "io"

type Transaction struct {
	Data []byte
}

func (*Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (*Transaction) DecodeBinary(r io.Reader) error {
	return nil
}
