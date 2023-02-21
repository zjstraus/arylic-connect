/*
arylic-connect, an API broker for Arylic Audio devices
Copyright (C) 2023  Zach Strauss

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package extWebsocket

import (
	"arylic-connect/rpcWrapper/websocketControl"
	"arylic-connect/transport/websocket"
	"context"
	"log"
	"sync"
	"time"
)

type ExternalWebsocketWrapper struct {
	HttpMediaCons map[string]*websocketControl.RPC
	OpLock        sync.RWMutex
}

func New() *ExternalWebsocketWrapper {
	return &ExternalWebsocketWrapper{
		HttpMediaCons: make(map[string]*websocketControl.RPC),
	}
}

func (wrapper *ExternalWebsocketWrapper) ConnectEndpoint(target string, name string) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*15)
	defer ctxCancel()

	transport, transportErr := websocket.New()
	if transportErr != nil {
		return transportErr
	}
	connectErr := transport.Connect(target)
	if connectErr != nil {
		return connectErr
	}
	rpc := websocketControl.New(transport)
	_, statusErr := rpc.GetStatus(ctx)
	if statusErr != nil {
		return statusErr
	}

	wrapper.OpLock.Lock()
	defer wrapper.OpLock.Unlock()

	existingEndpoint, hasEndpoint := wrapper.HttpMediaCons[target]
	if hasEndpoint {
		closeErr := existingEndpoint.Close()
		if closeErr != nil {
			log.Printf("Error closing existing endpoint connection for %s: %s\n", name, closeErr)
		}
		delete(wrapper.HttpMediaCons, name)
	}

	rpc.GetStatus(ctx)

	wrapper.HttpMediaCons[name] = rpc

	return nil
}

type EndpointInfo struct {
	Name   string
	Target string
}

func (wrapper *ExternalWebsocketWrapper) ConnectedEndpoints() []EndpointInfo {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()
	endpoints := make([]EndpointInfo, 0)

	for name, endpoint := range wrapper.HttpMediaCons {
		endpoints = append(endpoints, EndpointInfo{
			Name:   name,
			Target: "ws://" + endpoint.TransportTarget() + "/",
		})
	}

	return endpoints
}
