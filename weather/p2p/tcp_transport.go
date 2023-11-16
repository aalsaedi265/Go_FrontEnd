package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct{
	conn net.Conn

	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool)*TCPPeer{
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

type TCPTransport struct{
	ListenAddress string
	listener net.Listener
	handShakeFunc HandShakeFunc

	mu sync.RWMutex
	peer map[net.Addr]Peer
}


func NewTCPTransport(listenAddr string)*TCPTransport{
	return &TCPTransport{
		handShakeFunc: func(any) error{return nil},
		ListenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept()error{
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil{
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop(){
	for{
		conn, err := t.listener.Accept()
		if err != nil{
			fmt.Printf("TCP accept error: %s\n", err)
		}
		
		go t.handleConn(conn)
	}
}
func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn, true)

	// if err := t.handShakeFunc(conn); err != nil{
	// }
	// for{
	// 	n, err
	// }
	fmt.Printf("new incoming connection %+v\n", peer)
}

