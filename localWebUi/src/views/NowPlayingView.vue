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
import {useSerialMediaApiTransport} from "@/stores/serialMediaTransport";

const wsStore = useWebsocketStore()
const wsApi = useExternalWebsocketApi()
const wsStatus = useExternalWebsocketApiStatus()
const serialMetadataStore = useSerialMediaApiMetadata()
const serialMediaStore = useSerialMediaApi()
const serialTransportStore = useSerialMediaApiTransport()
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

async function transportNext() {
  await serialTransportStore.triggerNext()
}

async function transportPrevious() {
  await serialTransportStore.triggerPrevious()
}

async function transportStop() {
  await serialTransportStore.triggerStop()
}

async function transportPlayPause() {
  await serialTransportStore.triggerPlayPause()
}

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
  white-space: nowrap;
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
  white-space: nowrap;
}

.btn {
  color: white;
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
      <v-card variant="tonal" class="d-flex flex-column" color="rgba(0,0,0,.2)" style="flex-grow: 1; padding: 0px 10px;">
        <div class="titleText overflow-hidden">
          {{ wsStatus.title }}
        </div>
        <div class="infoText">
          {{ wsStatus.artist }}
        </div>
        <div class="infoText">
          {{ wsStatus.album }}
        </div>
        <div class="infoText">
          <div class="d-flex align-center" style="gap: 5px">
            <div>
              {{ new Date(wsStatus.elapsed * 1000).toISOString().slice(14, 19) }}
            </div>
            <v-progress-linear style="flex-grow: 1;" :max="wsStatus.duration" :model-value="wsStatus.elapsed">
            </v-progress-linear>
            <div>
              {{ new Date(wsStatus.duration * 1000).toISOString().slice(14, 19) }}
            </div>
          </div>
        </div>

        <div class="d-flex" style="gap: 10px; flex-grow: 1">

          <v-btn variant="tonal" rounded="false" @click="transportPrevious" style="flex-grow: 1; height: 100%">
            <v-icon color="grey-lighten-1" size="68"  icon="mdi-skip-previous"></v-icon>
          </v-btn>
<!--          <v-btn variant="tonal" rounded="false" @click="transportStop" style="flex-grow: 1; height: 100%">-->
<!--            <v-icon color="white" size="72"  icon="mdi-stop"></v-icon>-->
<!--          </v-btn>-->
          <v-btn variant="tonal" rounded="false" @click="transportPlayPause" style="flex-grow: 1.5; height: 100%;">
            <v-icon color="white" size="72"  icon="mdi-play-pause"></v-icon>
          </v-btn>
          <v-btn variant="tonal" rounded="false" @click="transportNext" style="flex-grow: 1; height: 100%">
            <v-icon color="grey-lighten-1" size="68" icon="mdi-skip-next"></v-icon>
          </v-btn>
        </div>
      </v-card>
    </div>
  </main>
</template>
