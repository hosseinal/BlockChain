package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
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

	hash types.Hash
}

func (b *Block) Hash() types.Hash {

	if !b.hash.IsZero() {
		return b.hash
	}

	buf := bytes.Buffer{}
	err := b.EncodeBinary(&buf)
	if err != nil {
		fmt.Println("Error encoding block for hash:", err)
		return types.Hash{}
	}

	b.hash = sha256.Sum256(buf.Bytes())

	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {
	err := b.Header.EncodeBinary(w)
	if err != nil {
		return err
	}
	// Dummy implementation for example purposes
	for _, tx := range b.Transactions {
		err := tx.EncodeBinary(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	err := b.Header.DecodeBinary(r)
	if err != nil {
		return err
	}
	// Dummy implementation for example purposes
	// Assuming we know how many transactions to read; here we just read 2 for example
	for i := range b.Transactions {
		err := b.Transactions[i].DecodeBinary(r)
		if err != nil {
			return err
		}
	}
	return nil
}
