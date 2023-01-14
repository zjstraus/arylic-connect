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

package httpmedia

import (
	"arylic-connect/rpcWrapper/httpControl"
	"arylic-connect/transport/http"
	"context"
	"log"
	"sync"
	"time"
)

type HttpMediaWrapper struct {
	HttpMediaCons map[string]*httpControl.RPC
	OpLock        sync.RWMutex
}

func New() *HttpMediaWrapper {
	return &HttpMediaWrapper{
		HttpMediaCons: make(map[string]*httpControl.RPC),
	}
}

func (wrapper *HttpMediaWrapper) ConnectEndpoint(target string) (string, error) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*15)
	defer ctxCancel()

	transport, _ := http.New()
	connectErr := transport.Connect(target)
	if connectErr != nil {
		return "", connectErr
	}
	rpc := httpControl.New(transport)
	status, statusErr := rpc.GetStatus(ctx)
	if statusErr != nil {
		return "", statusErr
	}

	wrapper.OpLock.Lock()
	defer wrapper.OpLock.Unlock()

	existingEndpoint, hasEndpoint := wrapper.HttpMediaCons[status.DeviceName]
	if hasEndpoint {
		closeErr := existingEndpoint.Close()
		if closeErr != nil {
			log.Printf("Error closing existing endpoint connection for %s: %s\n", status.DeviceName, closeErr)
		}
		delete(wrapper.HttpMediaCons, status.DeviceName)
	}

	wrapper.HttpMediaCons[status.DeviceName] = rpc

	return status.DeviceName, nil
}

type EndpointInfo struct {
	Name   string
	Target string
}

func (wrapper *HttpMediaWrapper) ConnectedEndpoints() []EndpointInfo {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()
	endpoints := make([]EndpointInfo, 0)

	for name, endpoint := range wrapper.HttpMediaCons {
		endpoints = append(endpoints, EndpointInfo{
			Name:   name,
			Target: endpoint.TransportTarget(),
		})
	}

	return endpoints
}
