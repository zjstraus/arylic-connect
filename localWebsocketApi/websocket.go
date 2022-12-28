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
	"net/http"
)

func RunWebsocketServer() error {
	serialMediaService := serialmedia.New()

	rpcServer := rpc.NewServer()
	serialMediaErr := rpcServer.RegisterName("serialmedia", serialMediaService)
	if serialMediaErr != nil {
		panic(serialMediaErr)
	}

	http.Handle("/ws", rpcServer.WebsocketHandler([]string{"*"}))
	return http.ListenAndServe(":8080", nil)
}
