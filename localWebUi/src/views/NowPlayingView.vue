<!--arylic-connect, an API broker for Arylic Audio devices-->
<!--Copyright (C) 2022  Zach Strauss-->

<!--This program is free software: you can redistribute it and/or modify-->
<!--it under the terms of the GNU General Public License as published by-->
<!--the Free Software Foundation, either version 3 of the License, or-->
<!--(at your option) any later version.-->

<!--This program is distributed in the hope that it will be useful,-->
<!--but WITHOUT ANY WARRANTY; without even the implied warranty of-->
<!--MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the-->
<!--GNU General Public License for more details.-->

<!--You should have received a copy of the GNU General Public License-->
<!--along with this program.  If not, see <https://www.gnu.org/licenses/>.-->

<script setup lang="ts">
import {useWebsocketStore} from "@/stores/baseWebsocket";
import {onMounted} from "vue";
import {useUIStateStore} from "@/stores/uiState";
import {useSerialMediaApiMetadata} from "@/stores/serialMediaMetadata";
import {useSerialMediaApi} from "@/stores/serialMediaApi";
import { useRoute } from 'vue-router'
import { ref, watch } from 'vue'
import {useExternalWebsocketApiStatus} from "@/stores/extWebsocketStatus";
import {useExternalWebsocketApi} from "@/stores/extWebsocketApi";

const wsStore = useWebsocketStore()
const wsApi = useExternalWebsocketApi()
const wsStatus = useExternalWebsocketApiStatus()
const serialMetadataStore = useSerialMediaApiMetadata()
const serialMediaStore = useSerialMediaApi()
const uiState = useUIStateStore()

onMounted(async () => {
  uiState.appBarActive = false
  await wsStore.connect()
  serialMediaStore.activePlayer = route.params.player as string
  wsApi.activePlayer = route.params.player as string
  await serialMetadataStore.subscribeMetadata()
  await wsStatus.pollStatus()
})

const route = useRoute()

watch(
    () => route.params.player,
    async newPlayer => {
      serialMediaStore.activePlayer = newPlayer as string
    }
)

watch(
    () => serialMetadataStore.title,
    async newTitle => {
      await wsStatus.pollStatus()
    }
)

setInterval(async () => {
  await wsStatus.pollStatus()
}, 10000)

</script>

<style>
/* Text */
.titleText {
  width: 100%;
  height: 80px;
  color: rgba(255, 255, 255, 1);
  font-family: sans-serif;
  font-style: normal;
  font-size: 48px;
  font-weight: 400;
  line-height: 1.2;
  letter-spacing: 0px;
  text-decoration: none;
  text-transform: none;
}
/* Text */
.infoText {
  width: 100%;
  height: 50px;
  color: rgba(255, 255, 255, 1);
  font-family: sans-serif;
  font-style: normal;
  font-size: 36px;
  font-weight: 400;
  line-height: 1.2;
  letter-spacing: 0px;
  text-decoration: none;
  text-transform: none;
}
</style>

<template>
  <main>
    <div :style="{backgroundImage:`url(${wsStatus.image})`, backgroundSize:'100% 100%', filter: 'blur(30px) brightness(50%)', position:'absolute', height:'100%', width:'100%'}">
    </div>

    <div class="d-flex" style="padding: 10px 20px; gap: 15px; z-index: 2;position: absolute; width: 100%" >
      <v-card width="300">
        <Transition>
        <v-img aspect-ratio="1" :src="wsStatus.image"></v-img>
        </Transition>
      </v-card>
      <v-card variant="tonal" class="d-flex flex-column" color="rgba(0,0,0,.2)" style="flex-grow: 1; padding: 0px 10px; backdrop-filter: blur()">
        <div class="titleText">
          {{ wsStatus.title }}
        </div>
        <div class="infoText">
          {{ wsStatus.artist }}
        </div>
        <div class="infoText">
          {{ wsStatus.album }}
        </div>
      </v-card>
      <v-card variant="tonal" width="100">Controls</v-card>
    </div>
  </main>
</template>
