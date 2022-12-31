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
    <v-card-title>Volume</v-card-title>
    <v-card-item>
    <v-form>
    <v-slider label="Volume" v-model="volume" min="0" max="1" step="0.01"></v-slider>
    <v-slider label="Balance" v-model="balance" min="-1" max="1" step="0.1"></v-slider>
    <v-switch label="Mute" v-model="mute" color="primary" inset></v-switch>
    </v-form>
    </v-card-item>
  </v-card>
</template>

<script setup lang="ts">
import {onMounted} from "vue";
import {storeToRefs} from "pinia";
import {useSerialMediaApiEq} from "@/stores/serialMediaEq";
import {useSerialMediaApi} from "@/stores/serialMediaApi";
import {useSerialMediaApiVolume} from "@/stores/serialMediaVolume";

const {pollAll} = useSerialMediaApiVolume()

const {volume, mute, balance} = storeToRefs(useSerialMediaApiVolume())

onMounted(async () => {
  await pollAll()
})

const topSerialMedia = useSerialMediaApi()
topSerialMedia.$subscribe(() => {
  pollAll()
})
</script>