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

package localWebsocketApi

import (
	"arylic-connect/localWebsocketApi/serialmedia"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/koron/go-ssdp"
	"log"
	"net/http"
	"net/url"
	"time"
)

type WebsocketManager struct {
	serialConnections *serialmedia.SerialMediaWrapper
}

func (manager *WebsocketManager) discoverSsdp() {
	ssdpList, ssdpErr := ssdp.Search(ssdp.All, 5, "")
	if ssdpErr != nil {
		panic(ssdpErr)
	}
	knownEndpoints := manager.serialConnections.ConnectedEndpoints()
	for _, service := range ssdpList {
		if service.Type != "urn:schemas-wiimu-com:service:PlayQueue:1" {
			continue
		}
		parsedUrl, urlErr := url.Parse(service.Location)
		if urlErr == nil {
			targetAddr := parsedUrl.Hostname() + ":8899"
			alreadyConnected := false
			for _, endpoint := range knownEndpoints {
				if endpoint.Target == targetAddr {
					alreadyConnected = true
					break
				}
			}
			if !alreadyConnected {
				log.Printf("Discovered potential device at %s\n", parsedUrl.Hostname())
				name, connectErr := manager.serialConnections.ConnectEndpoint(targetAddr)
				if connectErr == nil {
					log.Printf("TCP connected to player %s\n", name)
				}
			}
		}
	}
}

func (manager *WebsocketManager) ssdpLoop() {
	ticker := time.NewTicker(time.Minute)
	manager.discoverSsdp()
	for {
		select {
		case <-ticker.C:
			manager.discoverSsdp()
		}
	}
}

func (manager *WebsocketManager) wsRpcLoop() error {
	rpcServer := rpc.NewServer()
	serialMediaErr := rpcServer.RegisterName("serialmedia", manager.serialConnections)
	if serialMediaErr != nil {
		panic(serialMediaErr)
	}

	http.Handle("/ws", rpcServer.WebsocketHandler([]string{"*"}))
	log.Println("Starting web server")
	return http.ListenAndServe(":8080", nil)
}

func RunWebsocketServer() error {
	manager := WebsocketManager{
		serialConnections: serialmedia.New(),
	}

	go manager.ssdpLoop()
	return manager.wsRpcLoop()
}
