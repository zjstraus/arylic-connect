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

interface metadataUpdateMessage {
    album: string
    artist: string
    skiplimit: number
    title: string
    vendor: string
}

export const useSerialMediaApiMetadata = defineStore("serialMediaApiMetadata", () => {
    const api = useSerialMediaApi()

    const album = ref("")
    const artist = ref("")
    const skiplimit = ref(0)
    const title = ref("")
    const vendor = ref("")

    async function subscribeMetadata() {
        await api.addSubscription("metadataChanges", (data: object) => {
            let cast = data as metadataUpdateMessage
            album.value = cast.album
            artist.value = cast.artist
            skiplimit.value = cast.skiplimit
            title.value = cast.title
            vendor.value = cast.vendor
        })
    }


    return {subscribeMetadata, album, artist, skiplimit, title, vendor}
})