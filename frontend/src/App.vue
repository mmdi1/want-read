
<template>
  <div class="main"  style="--wails-draggable: drag">
    <Index v-bind:mode="mode" v-show="mode.screen == 'main'" />
    <Reading v-bind:mode="mode" v-show="mode.screen == 'read'"></Reading>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, reactive } from "vue";
import Index from "./pages/index.vue";
import Reading from "./pages/reading.vue";
import { Ws } from "./common/websocket";
window.addEventListener("resize", () => {
  console.log(window.innerWidth, window.innerHeight);
});
let mode: any = reactive({
  screen: "main",
  content: "",
});
const updateContent = (str: string) => {
  mode.content = str;
};

let websocket = new Ws("ws://127.0.0.1:8899/ws");
websocket.link.onmessage = (event) => {
  let obj = JSON.parse(event.data);
  updateContent(obj.data);
  console.log("=============>",obj);
};
</script>

<style>

</style>
