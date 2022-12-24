package tcp

import (
	"arylic-connect/transport"
	"net"
	"sync"
	"time"
)

type Transport struct {
	conn net.Conn

	listenerCloser     chan int
	requestLocker      sync.Mutex
	persistentRequests map[string][]chan<- []byte
	oneshotRequests    map[string][]chan<- []byte

	outgoingQueue chan string
}

func New() (*Transport, error) {
	return &Transport{
		persistentRequests: make(map[string][]chan<- []byte),
		oneshotRequests:    make(map[string][]chan<- []byte),
	}, nil
}

func (t *Transport) Connect(target string) error {
	t.Close()

	t.listenerCloser = make(chan int)
	t.outgoingQueue = make(chan string)

	conn, err := net.DialTimeout("tcp", target, 10*time.Second)
	if err != nil {
		return err
	}
	t.conn = conn
	go t.asyncReadLoop()
	go t.asyncWriteLoop()
	return nil
}

func (t *Transport) Close() error {
	if t.conn != nil {
		return t.conn.Close()
	}
	if t.listenerCloser != nil {
		close(t.listenerCloser)
	}

	return nil
}

func (t *Transport) Flavor() transport.InterfaceFlavor {
	return transport.Flavor_TCP
}
