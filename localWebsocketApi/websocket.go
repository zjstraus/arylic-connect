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
	"arylic-connect/localWebsocketApi/httpmedia"
	"arylic-connect/localWebsocketApi/serialmedia"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/koron/go-ssdp"
	"log"
	"net/http"
	"net/url"
	"time"
)

type WebsocketManager struct {
	serialConnections *serialmedia.SerialMediaWrapper
	httpConnections   *httpmedia.HttpMediaWrapper
}

func (manager *WebsocketManager) discoverSsdp() {
	ssdpList, ssdpErr := ssdp.Search(ssdp.All, 5, "")
	if ssdpErr != nil {
		panic(ssdpErr)
	}
	knownSerialEndpoints := manager.serialConnections.ConnectedEndpoints()
	knownHttpEndpoints := manager.httpConnections.ConnectedEndpoints()
	for _, service := range ssdpList {
		if service.Type != "urn:schemas-wiimu-com:service:PlayQueue:1" {
			continue
		}
		parsedUrl, urlErr := url.Parse(service.Location)
		if urlErr == nil {
			serialTarget := parsedUrl.Hostname() + ":8899"
			httpTarget := fmt.Sprintf("http://%s/httpapi.asp", parsedUrl.Hostname())
			serialConnected := false
			httpConnected := false
			for _, endpoint := range knownSerialEndpoints {
				if endpoint.Target == serialTarget {
					serialConnected = true
					break
				}
			}
			for _, endpoint := range knownHttpEndpoints {
				if endpoint.Target == httpTarget {
					httpConnected = true
					break
				}
			}
			if !serialConnected {
				log.Printf("Discovered potential device at %s\n", parsedUrl.Hostname())
				name, connectErr := manager.serialConnections.ConnectEndpoint(serialTarget)
				if connectErr == nil {
					log.Printf("TCP connected to player %s\n", name)
				}
			}
			if !httpConnected {
				log.Printf("Discovered potential device at %s\n", parsedUrl.Hostname())
				name, connectErr := manager.httpConnections.ConnectEndpoint(httpTarget)
				if connectErr == nil {
					log.Printf("HTTP connected to player %s\n", name)
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
	httpMediaErr := rpcServer.RegisterName("httpmedia", manager.httpConnections)
	if httpMediaErr != nil {
		panic(httpMediaErr)
	}

	uiDist := http.FileServer(http.Dir("localWebUi/dist"))
	http.Handle("/", uiDist)
	http.Handle("/ws", rpcServer.WebsocketHandler([]string{"*"}))
	log.Println("Starting web server")
	return http.ListenAndServe(":8080", nil)
}

func RunWebsocketServer() error {
	manager := WebsocketManager{
		serialConnections: serialmedia.New(),
		httpConnections:   httpmedia.New(),
	}

	go manager.ssdpLoop()
	return manager.wsRpcLoop()
}
