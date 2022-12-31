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
import {onMounted, ref, computed} from "vue";
import {useWebsocketStore} from "@/stores/baseWebsocket";
import {useSerialMediaApi} from "@/stores/serialMediaApi";

const theme = ref('dark')
const drawer = ref(false)
const wsStore = useWebsocketStore()
const serialMediaApi = useSerialMediaApi()

onMounted(async () => {
  await wsStore.connect()
  await wsStore.refreshEndpoints()
})

const endpointNames = computed(() => {
  let formatted = []
  for (const endpoint of wsStore.endpoints) {
    formatted.push(endpoint.Name)
  }
  formatted.sort()
  return formatted
})

function toggleTheme() {
  theme.value = theme.value === 'light' ? 'dark' : 'light'
}

const activePlayer = ref("")
function setActivePlayer() {
  serialMediaApi.activePlayer = activePlayer.value
  serialMediaApi.getEndpointVersion()
}
</script>

<template>
  <v-app :theme="theme">
    <v-app-bar>
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer">
      </v-app-bar-nav-icon>

      <v-app-bar-title>Arylic-Connect</v-app-bar-title>

      <v-spacer></v-spacer>

      <v-btn :prepend-icon="theme === 'light' ? 'mdi-weather-sunny' : 'mdi-weather-night'"
             @click="toggleTheme">
      </v-btn>

    </v-app-bar>

    <v-navigation-drawer v-model="drawer">
      <v-list-item>
        <v-combobox
            label="player"
            v-model="activePlayer"
            @change="setActivePlayer"
            :items="endpointNames"
        ></v-combobox>
      </v-list-item>
    </v-navigation-drawer>

    <v-main><router-view></router-view></v-main>
  </v-app>
</template>
