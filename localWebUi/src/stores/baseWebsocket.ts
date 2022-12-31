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

import {defineStore} from "pinia";
import {RpcWebSocketClient} from "rpc-websocket-client";

interface endpointIdentifier {
    Name: string,
    Target: string
}

interface websocketStoreState {
    client: RpcWebSocketClient,
    endpoints: endpointIdentifier[]
}

export const useWebsocketStore = defineStore('websocket',  {
    state: (): websocketStoreState => ({
        client: new RpcWebSocketClient(),
        endpoints: []
    }),
    getters: {
    },
    actions: {
        async connect () {
            if (this.client.ws != undefined) {
                if (this.client.ws.readyState == 1) {
                    return
                }
            }
            await this.client.connect("ws://localhost:8080/ws")
        },
        async refreshEndpoints() {
            this.endpoints = await this.client.call("serialmedia_connectedEndpoints") as endpointIdentifier[]
        }
    }
})