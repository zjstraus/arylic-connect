package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"arylic-connect/transport/tcp"
	"sync"
)

type SerialMediaWrapper struct {
	SerialMediaCons map[string]*serialMediaControl.RPC
	OpLock          sync.RWMutex
}

func New() *SerialMediaWrapper {
	return &SerialMediaWrapper{
		SerialMediaCons: make(map[string]*serialMediaControl.RPC),
	}
}

func (wrapper *SerialMediaWrapper) ConnectEndpoint(target string) error {
	wrapper.OpLock.Lock()
	existingEndpoint, hasEndpoint := wrapper.SerialMediaCons[target]
	if hasEndpoint {
		existingEndpoint.Close()
		delete(wrapper.SerialMediaCons, target)
	}
	wrapper.OpLock.Unlock()

	transport, _ := tcp.New()
	connectErr := transport.Connect(target)
	if connectErr != nil {
		return connectErr
	}
	rpc := serialMediaControl.New(transport)

	wrapper.OpLock.Lock()
	wrapper.SerialMediaCons[target] = rpc
	wrapper.OpLock.Unlock()

	return nil
}
