package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct{
	ListenAddress string
	listener net.Listener

	mu sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string)*TCPTransport{
	return &TCPTransport{
		ListenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept()error{
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil{
		return err
	}
}

func (t *TCPTransport) acceptLoop(){
	
}