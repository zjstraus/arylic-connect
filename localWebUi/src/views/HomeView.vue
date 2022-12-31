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
import SerialMediaApiVersion from "@/components/homeCards/serialMediaApiVersion.vue";
import SerialMediaApiEq from "@/components/homeCards/serialMediaApiEq.vue";
import SerialMediaApiVolume from "@/components/homeCards/serialMediaApiVolume.vue";

const wsStore = useWebsocketStore()
onMounted(async () => {
  await wsStore.connect()
})

</script>

<template>
  <main>
    <v-container fluid>
      <v-row>
        <v-col cols="4">
          <serial-media-api-eq></serial-media-api-eq>
        </v-col>
      <v-col cols="3">
        <serial-media-api-volume></serial-media-api-volume>
      </v-col>
        <v-col cols="3">
           <serial-media-api-version></serial-media-api-version>
        </v-col>
      </v-row>
    </v-container>
  </main>
</template>
