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

import {defineStore} from "pinia";
import {RpcWebSocketClient} from "rpc-websocket-client";

interface endpointIdentifier {
    Name: string,
    Target: string
}

interface websocketStoreState {
    client: RpcWebSocketClient,
    endpoints: endpointIdentifier[]
    subsciptionCallbacks: Map<string, (data: object) => void>
}

interface subscriptionReturn {
    result: string
}

interface subscriptionUpdate {
    result: any
    subscription: string
}

interface subscriptionUpdateWrapper {
    params: subscriptionUpdate
}

export const useWebsocketStore = defineStore('websocket',  {
    state: (): websocketStoreState => ({
        client: new RpcWebSocketClient(),
        endpoints: [],
        subsciptionCallbacks: new Map<string, (data: object) => void>()
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
            if (import.meta.env.DEV) {
                await this.client.connect(`ws://localhost:8080/ws`)
            } else {
                await this.client.connect(`ws://${window.location.host}/ws`)
            }
            this.client.onNotification.push((data: object) => {
                let cast = data as subscriptionUpdateWrapper
                let cb = this.subsciptionCallbacks.get(cast.params.subscription)
                if (cb) {
                    cb(cast.params.result)
                }
            })
            this.client.onRequest.push((data: object) => {
                console.log(data)
            })
            //this.client.listenMessages()
        },
        async refreshEndpoints() {
            this.endpoints = await this.client.call("serialmedia_connectedEndpoints") as endpointIdentifier[]
        },
        async addSubscription(method: string, params: any, cb: (data: object) => void) {
            await this.connect()
            let result = await this.client.call(method, params) as string
            this.subsciptionCallbacks.set(result, cb)
        }
    }
})