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


export const useSerialMediaApiEq = defineStore("serialMediaApiNetwork", () => {
    const api = useSerialMediaApi()

    const internet = ref(false)
    async function sendInternet() {
        internet.value = await api.makeCall("setInternet", [internet.value]) as boolean
    }
    async function pollInternet() {
        internet.value = await api.makeCall("getInternet", []) as boolean
    }

    const ethernet = ref(false)
    async function sendEthernet() {
        ethernet.value = await api.makeCall("setEthernet", [ethernet.value]) as boolean
    }
    async function pollEthernet() {
        ethernet.value = await api.makeCall("getEthernet", []) as boolean
    }

    const wifi = ref(false)
    async function sendWifi() {
        wifi.value = await api.makeCall("setWifi", [wifi.value]) as boolean
    }
    async function pollWifi() {
        wifi.value = await api.makeCall("getWifi", []) as boolean
    }

    const bluetooth = ref(false)
    async function sendBluetooth() {
        bluetooth.value = await api.makeCall("setBluetooth", [bluetooth.value]) as boolean
    }
    async function pollBluetooth() {
        bluetooth.value = await api.makeCall("getBluetooth", []) as boolean
    }

    const wifiPlayback = ref(false)
    async function pollWifiPlayback() {
        wifiPlayback.value = await api.makeCall("getWifiPlayback", []) as boolean
    }

    async function pollAll() {
        return Promise.all([
            pollInternet(),
            pollWifi(),
            pollEthernet(),
            pollBluetooth(),
            pollWifiPlayback()
        ])
    }

    async function sendAll() {
        return Promise.all([
            sendInternet(),
            sendWifi(),
            sendEthernet(),
            sendBluetooth()
        ])
    }

    return {internet, wifi, ethernet, bluetooth, wifiPlayback,
        pollInternet, pollWifi, pollEthernet, pollBluetooth, pollWifiPlayback,
        sendInternet, sendWifi, sendEthernet, sendBluetooth,
        sendAll, pollAll}
})