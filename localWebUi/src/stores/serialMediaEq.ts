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


export const useSerialMediaApiEq = defineStore("serialMediaApiEq", () => {
    const api = useSerialMediaApi()

    const bass = ref(0)
    watch(bass, async (level) => {
        bass.value = await api.makeCall("setBass", [level]) as number
    })
    async function pollBass() {
        bass.value = await api.makeCall("getBass", []) as number
    }

    const treble = ref(0)
    watch(treble, async (level) => {
        treble.value = await api.makeCall("setTreble", [level]) as number
    })
    async function pollTreble() {
        treble.value = await api.makeCall("getTreble", []) as number
    }

    const virtualBass = ref(false)
    watch(virtualBass, async(enabled) => {
        virtualBass.value = await api.makeCall("setVirtualBass", [enabled]) as boolean
    })
    async function pollVirtualBass() {
        virtualBass.value = await api.makeCall("getVirtualBass", []) as boolean
    }

    async function pollAll() {
        return Promise.all([
            pollTreble(),
            pollBass(),
            pollVirtualBass()
        ])
    }

    return {bass, treble, virtualBass, pollBass, pollTreble, pollVirtualBass, pollAll}
})