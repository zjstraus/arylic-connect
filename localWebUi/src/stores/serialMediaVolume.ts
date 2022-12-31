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

export const useSerialMediaApiVolume = defineStore("serialMediaApiVolume", () => {
    const api = useSerialMediaApi()

    const volume = ref(0)
    watch(volume, async (level) => {
        volume.value = await api.makeCall("setVolume", [level]) as number
    })
    async function pollVolume() {
        volume.value = await api.makeCall("getVolume", []) as number
    }

    const maxVolume = ref(1)
    watch(maxVolume, async (level) => {
        maxVolume.value = await api.makeCall("setMaxVolume", [level]) as number
    })
    async function pollMaxVolume() {
        maxVolume.value = await api.makeCall("getMaxVolume", []) as number
    }

    const fixedVolume = ref(false)
    watch(fixedVolume, async(enabled) => {
        fixedVolume.value = await api.makeCall("setFixedVolume", [enabled]) as boolean
    })
    async function pollFixedVolume() {
        fixedVolume.value = await api.makeCall("getFixedVolume", []) as boolean
    }

    const mute = ref(false)
    watch(mute, async(enabled) => {
        mute.value = await api.makeCall("setMute", [enabled]) as boolean
    })
    async function pollMute() {
        mute.value = await api.makeCall("getMute", []) as boolean
    }

    const balance = ref(0)
    watch(balance, async (setting) => {
        balance.value = await api.makeCall("setBalance", [setting]) as number
    })
    async function pollBalance() {
        balance.value = await api.makeCall("getBalance", []) as number
    }

    async function pollAll() {
        return Promise.all([
            pollVolume(),
            pollMaxVolume(),
            pollMute(),
            pollBalance(),
            pollFixedVolume()
        ])
    }

    return {volume, maxVolume, balance, fixedVolume, mute, pollVolume, pollMaxVolume, pollBalance, pollMute, pollFixedVolume,  pollAll}
})