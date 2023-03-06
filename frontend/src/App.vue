
<template>
  <div class="main" style="--wails-draggable: drag">
    <n-config-provider :theme-overrides="themeOverrides">
      <div :class="mode.screen == 'main' ? 'main-bg' : 'read-bg'">
        <Index v-bind:mode="mode" v-show="mode.screen == 'main'" />
        <Reading
          id="readDiv"
          v-bind:mode="mode"
          v-show="mode.screen == 'read'"
        ></Reading>
      </div>
    </n-config-provider>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, reactive, ref } from "vue";
import Index from "./pages/index.vue";
import Reading from "./pages/reading.vue";
import { Ws } from "./common/websocket";
import { Environment, WindowSetSize } from "../wailsjs/runtime/runtime";
import { SetReadPanel } from "../wailsjs/go/setting/App";
import { NConfigProvider, GlobalThemeOverrides } from "naive-ui";

const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: "#FF0000",
  },
  Button: {
    textColor: "#FF0000",
  },
};

window.addEventListener("resize", async () => {
  if (mode.screen == "read") {
    console.log(window.innerWidth, window.innerHeight);
    await SetReadPanel(window.innerWidth, window.innerHeight);
  }
});
let mode: any = reactive({
  screen: "main",
  bookshelf: [],
  content: "",
  showSize: 100,
  fontSize: 12,
  fontColor: "rgb(220,220,220)",
  prevGroup: [],
  nextGroup: [],
  hideGroup: [],
  show: true,
});
let show = ref(true);
const updateContent = (str: string) => {
  mode.content = str;
};
mode.screen = "main";
let websocket = new Ws("ws://127.0.0.1:8899/ws");
websocket.link.onmessage = async (event) => {
  let obj = JSON.parse(event.data);
  console.log("=============>", obj);
  switch (obj.id) {
    case 0: //上一页
    case 1: //下一页
      updateContent(obj.data);
      break;
    case 2: //隐藏
      mode.show = !mode.show;
      break;
    case 3: //配置界面
      mode.screen = "main";
      break;
    case 4: //退出
      await Environment();
    case 5: //更新配置
      mode.fontSize = obj.data.font_size;
      mode.fontColor = obj.data.font_color;
      if (obj.data.read_width && obj.data.read_height) {
        await WindowSetSize(obj.data.read_width, obj.data.read_height);
      }
      break;
    default:
      break;
  }
};
</script>

<style>
.main-bg {
  background-color: white;
}
</style>
