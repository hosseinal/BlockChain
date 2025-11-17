package core

import (
	"encoding/binary"
	"io"

	"github.com/hosseinal/BlockChain/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp int64
	Heght     uint64
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	// Dummy implementation for example purposes
	err := binary.Write(w, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, &h.PrevBlock)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, &h.Timestamp)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, &h.Heght)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, &h.Nonce)
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	// Dummy implementation for example purposes
	err := binary.Read(r, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, &h.PrevBlock)
	if err != nil {
		return err
	}

	var ts int64
	err = binary.Read(r, binary.LittleEndian, &ts)
	if err != nil {
		return err
	}
	h.Timestamp = ts

	err = binary.Read(r, binary.LittleEndian, &h.Heght)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, &h.Nonce)
	if err != nil {
		return err
	}

	return nil
}

type Block struct {
	Header       Header
	Transactions []Transaction
}
