/*
arylic-connect, an API broker for Arylic Audio devices
Copyright (C) 2022  Zach Strauss

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

type EndpointInfo struct {
	Name   string
	Target string
}

func (wrapper *SerialMediaWrapper) ConnectedEndpoints() []EndpointInfo {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()
	endpoints := make([]EndpointInfo, 0)

	for name, endpoint := range wrapper.SerialMediaCons {
		endpoints = append(endpoints, EndpointInfo{
			Name:   name,
			Target: endpoint.TransportTarget(),
		})
	}

	return endpoints
}
