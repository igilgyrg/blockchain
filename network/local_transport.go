package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	lock      sync.RWMutex
	consumeCh chan RPC
	addr      NetAddr
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (l *LocalTransport) Consume() <-chan RPC {
	return l.consumeCh
}

func (l *LocalTransport) Connect(transport Transport) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.peers[transport.Addr()] = transport.(*LocalTransport)

	return nil
}

func (l *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	l.lock.RLock()
	defer l.lock.RUnlock()

	peer, ok := l.peers[to]
	if !ok {
		return fmt.Errorf("could not send message to %s", to)
	}

	peer.consumeCh <- RPC{
		From:    l.addr,
		Payload: payload,
	}

	return nil
}

func (l *LocalTransport) Addr() NetAddr {
	return l.addr
}
