package network

import "sync"

type LocalTransport struct {
	addr      NetAddr
	consumeSh chan RPC
	lock      sync.Mutex
	peers     map[NetAddr]*LocalTransport
}

func (lt *LocalTransport) Addr() NetAddr {
	return lt.addr
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeSh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (lt *LocalTransport) Consume() <-chan RPC {
	return lt.consumeSh
}

func (lt *LocalTransport) Connect(tr Transport) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()
	lt.peers[tr.Addr()] = tr.(*LocalTransport)
	return nil
}

func (lt *LocalTransport) SendMessage(addr NetAddr, payload []byte) error {
	lt.lock.Lock()
	peer, ok := lt.peers[addr]
	lt.lock.Unlock()
	if !ok {
		return nil
	}
	rpc := RPC{
		From:    lt.addr,
		Payload: payload,
	}
	peer.consumeSh <- rpc
	return nil
}
