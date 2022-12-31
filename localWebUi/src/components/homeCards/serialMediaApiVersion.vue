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

<template>
  <v-card>
    <v-card-title>Serial API Connection</v-card-title>
    <v-list lines="one">
      <v-list-item :title="endpointVersion.API" subtitle="API Version"></v-list-item>
      <v-list-item :title="endpointVersion.Firmware" subtitle="Firmware"></v-list-item>
      <v-list-item :title="endpointVersion.Git" subtitle="Git Revision"></v-list-item>
    </v-list>
  </v-card>
</template>

<script setup lang="ts">
import {onMounted} from "vue";
import {useSerialMediaApi} from "@/stores/serialMediaApi";
import {storeToRefs} from "pinia";

const {getEndpointVersion} = useSerialMediaApi()

const { endpointVersion } = storeToRefs(useSerialMediaApi())
onMounted(async () => {
  await getEndpointVersion()
})
</script>