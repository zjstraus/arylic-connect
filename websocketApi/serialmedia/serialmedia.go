package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"arylic-connect/transport/tcp"
	"context"
	"sync"
	"time"
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

func (wrapper *SerialMediaWrapper) ConnectEndpoint(target string) (string, error) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*15)
	defer ctxCancel()

	transport, _ := tcp.New()
	connectErr := transport.Connect(target)
	if connectErr != nil {
		return "", connectErr
	}
	rpc := serialMediaControl.New(transport)
	name, nameErr := rpc.GetName(ctx)
	if nameErr != nil {
		return "", nameErr
	}

	wrapper.OpLock.Lock()
	defer wrapper.OpLock.Unlock()

	existingEndpoint, hasEndpoint := wrapper.SerialMediaCons[name]
	if hasEndpoint {
		existingEndpoint.Close()
		delete(wrapper.SerialMediaCons, name)
	}

	wrapper.SerialMediaCons[name] = rpc

	return name, nil
}
