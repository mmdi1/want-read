<template>
  <div
    id="reading"
    :style="{ color: mode.fontColor, fontSize: mode.fontSize + 'px' }"
    @dblclick="openSetting"
  >
    {{ props.mode.show ? props.mode.content : "" }}
  </div>
</template>

<script lang="ts" setup>
import { WindowSetSize } from "../../wailsjs/runtime/runtime";
let props = defineProps<{
  mode: any;
}>();
const openSetting = async () => {
  props.mode.screen = "main";
  await WindowSetSize(600, 396);
};
</script>

<style>
#reading {
  color: gainsboro;
}
.reading {
  word-break: break-all;
}
</style>