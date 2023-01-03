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


export const useSerialMediaApiSettings = defineStore("serialMediaApiSettings", () => {
    const api = useSerialMediaApi()

    const led = ref(false)
    async function sendLed() {
        led.value = await api.makeCall("setLED", [led.value]) as boolean
    }
    async function pollLed() {
        led.value = await api.makeCall("getLED", []) as boolean
    }

    const beep = ref(false)
    async function sendBeep() {
        beep.value = await api.makeCall("setBeep", [beep.value]) as boolean
    }
    async function pollBeep() {
        beep.value = await api.makeCall("getBeep", []) as boolean
    }

    const name = ref("")
    async function sendName() {
        name.value = await api.makeCall("setName", [name.value]) as string
    }
    async function pollName() {
        name.value = await api.makeCall("getName", []) as string
    }

    const voicePrompt = ref(false)
    async function sendVoicePrompt() {
        voicePrompt.value = await api.makeCall("setVoicePrompt", [voicePrompt.value]) as boolean
    }
    async function pollVoicePrompt() {
        voicePrompt.value = await api.makeCall("getVoicePrompt", []) as boolean
    }

    async function pollAll() {
        return Promise.all([
            pollLed(),
            pollBeep(),
            pollName(),
            pollVoicePrompt(),
        ])
    }

    async function sendAll() {
        return Promise.all([
            sendLed(),
            sendBeep(),
            sendName(),
            sendVoicePrompt()
        ])
    }

    return {led, beep, name, voicePrompt,
        pollLed, pollBeep, pollName, pollVoicePrompt,
        sendLed, sendBeep, sendName, sendVoicePrompt,
        sendAll, pollAll}
})