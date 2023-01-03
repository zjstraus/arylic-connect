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
import {ref, watch} from "vue";
import {useSerialMediaApi} from "@/stores/serialMediaApi";


export enum Source {
    Network = "Network",
    USB = "USB",
    USBDAC = "USBDAC",
    Line1 = "Line-In",
    Line2 = "Line-In2",
    Bluetooth = "Bluetooth",
    Optical = "Optical",
    Coax = "Coax",
    I2S = "I2S",
    HDMI = "HDMI",
    None = "None",
    Unknown = "Unknown"
}

export const useSerialMediaApiSources = defineStore("serialMediaApiSources", () => {
    const api = useSerialMediaApi()

    const active = ref(Source.Unknown)
    async function sendActive() {
        active.value = await api.makeCall("setSource", [active.value]) as Source
    }
    async function pollActive() {
        active.value = await api.makeCall("getSource", []) as Source
    }

    const poweron = ref(Source.Unknown)
    async function sendPoweron() {
        poweron.value = await api.makeCall("setDefaultSource", [poweron.value]) as Source
    }
    async function pollPoweron() {
        poweron.value = await api.makeCall("getDefaultSource", []) as Source
    }

    const autoswitch = ref(false)
    async function sendAutoswitch() {
        autoswitch.value = await api.makeCall("setInputAutoswitch", [autoswitch.value]) as boolean
    }
    async function pollAutoswitch() {
        autoswitch.value = await api.makeCall("getInputAutoswitch", []) as boolean
    }

    async function pollAll() {
        return Promise.all([
            pollActive(),
            pollPoweron(),
            pollAutoswitch(),
        ])
    }

    async function sendAll() {
        return Promise.all([
            sendActive(),
            sendPoweron(),
            sendAutoswitch(),
        ])
    }

    return {active, poweron, autoswitch,
        pollActive, pollPoweron, pollAutoswitch,
        sendActive, sendPoweron, sendAutoswitch,
        sendAll, pollAll}
})