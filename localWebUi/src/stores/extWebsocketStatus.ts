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
import {useSerialMediaApi} from "@/stores/serialMediaApi";
import {ref, watch} from "vue";
import {ca} from "vuetify/locale";
import {useExternalWebsocketApi} from "@/stores/extWebsocketApi";

interface statusUpdateMessage {
    Input: string
    Source: string
    State: string
    Index: number
    Mode: string
    Elapsed: number
    Duration: number
    Title: string
    Artist: string
    Album: string
    Image: string
    Volume: number
}

export const useExternalWebsocketApiStatus = defineStore("externalWebsocketApiStatus", () => {
    const api = useExternalWebsocketApi()

    const input = ref("")
    const source = ref("")
    const state = ref("")
    const index = ref(0)
    const mode = ref("")
    const elapsed = ref(0)
    const duration = ref(0)
    const title = ref("")
    const album = ref("")
    const image = ref("")
    const volume = ref(0)
    const artist = ref("")

    async function pollStatus() {
        let msg = await api.makeCall("getStatus", []) as statusUpdateMessage
        input.value = msg.Input
        source.value = msg.Source
        state.value = msg.State
        index.value = msg.Index
        mode.value = msg.Mode
        elapsed.value = msg.Elapsed
        duration.value = msg.Duration
        title.value = msg.Title
        album.value = msg.Album
        image.value = msg.Image
        volume.value = msg.Volume
        artist.value = msg.Artist
    }
    setInterval(() => {
        if (elapsed.value < duration.value) {
            elapsed.value += 1
        }
    }, 1000)

    // async function subscribeStatus() {
    //     await api.addSubscription("metadataChanges", (data: object) => {
    //         let cast = data as metadataUpdateMessage
    //         album.value = cast.album
    //         artist.value = cast.artist
    //         skiplimit.value = cast.skiplimit
    //         title.value = cast.title
    //         vendor.value = cast.vendor
    //     })
    // }


    return {pollStatus, input, source, state, index, mode, elapsed, duration, title, album, image, volume, artist}
})