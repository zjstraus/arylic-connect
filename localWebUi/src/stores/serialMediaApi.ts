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
import {useWebsocketStore} from "@/stores/baseWebsocket";
import {ref} from "vue";
import {throttle} from "lodash";

interface serialmediaEndpointVersion {
    Git: string,
    Firmware: string,
    API: string
}


export const useSerialMediaApi = defineStore("serialMediaApi", () => {
    const ws = useWebsocketStore()

    const activePlayer = ref("")
    const endpointVersion = ref({
        Git: "",
        Firmware: "",
        API: ""
    } as serialmediaEndpointVersion)

    async function getEndpointVersion() {
        await ws.connect()
        endpointVersion.value = await ws.client.call("serialmedia_getVersion", [activePlayer.value]) as serialmediaEndpointVersion
    }

    const throttledCalls: { [key: string]: Function } = {}

    async function makeCall(submethod: string, params: any[]) {
        await ws.connect()
        if (! (submethod in throttledCalls)) {
            throttledCalls[submethod] = throttle((params: any[]) => {
                return ws.client.call("serialmedia_" + submethod, params)
            }, 200)
        }
        params.unshift(activePlayer.value)
        return throttledCalls[submethod](params)
        //return ws.client.call("serialmedia_" + submethod, params)
    }



    return {activePlayer, endpointVersion, getEndpointVersion, makeCall}
})