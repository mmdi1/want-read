<template>
  <div class="setting">
    <div>
      <n-divider class="title font12" title-placement="right"> 设置 </n-divider>
      <n-row>
        <n-col :span="12" class="key"> 上一页 </n-col>
        <n-col :span="12">
          <n-input
            :value="prevIpt"
            type="text"
            placeholder="Alt + ,"
            size="small"
            :on-focus="prevPageFocus"
            :on-blur="loseBlur"
          />
        </n-col>
      </n-row>
      <br />
      <n-row>
        <n-col :span="12" class="key"> 下一页 </n-col>
        <n-col :span="12">
          <n-input
            :value="nextIpt"
            type="text"
            placeholder="Alt + ."
            size="small"
            :on-focus="nextPageFocus"
            :on-blur="loseBlur"
        /></n-col>
      </n-row>
      <br />
      <n-row>
        <n-col :span="12" class="key"> 隐藏 </n-col>
        <n-col :span="12">
          <n-input
            :value="hideIpt"
            type="text"
            placeholder="Alt + M"
            size="small"
            :on-focus="hideWindowsFocus"
            :on-blur="loseBlur"
        /></n-col>
      </n-row>
      <br />
      <n-row>
        <n-col :span="12" class="key"> 每页字数 </n-col>
        <n-col :span="12">
          <n-input-number v-model:value="showSize" :min="10" size="small"
        /></n-col>
      </n-row>
      <br />
      <n-row>
        <n-col :span="12" class="key"> 字体大小 </n-col>
        <n-col :span="12">
          <n-input-number v-model:value="fontSize" :min="12" size="small"
        /></n-col>
      </n-row>
      <br />
      <n-row>
        <n-col :span="12" class="key"> 字体颜色 </n-col>
        <n-col :span="12">
          <n-color-picker size="small" v-model:value="fontColor" />
        </n-col>
      </n-row>
      <br />
      <n-row>
        <n-col :span="12" class="key"> 操作 </n-col>
        <n-col :span="12">
          <n-space justify="end">
            <n-button
              strong
              secondary
              type="warning"
              size="small"
              @click="clear"
            >
              初始
            </n-button>
            <n-button
              strong
              secondary
              type="info"
              size="small"
              @click="saveSetting"
            >
              保存
            </n-button>
          </n-space>
        </n-col>
      </n-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from "vue";
import Keys from "../assets/json/keys.json";
import { InitSetting, SaveSetting } from "../../wailsjs/go/setting/App";
let fontSize = ref(12);
let showSize = ref(50);
let fontColor = ref("rgb(220,220,220)");
let prevPageKey: string[] = [];
let nextPageKey: string[] = [];
let hideWindowsKey: string[] = [];
let prevIpt = ref(""),
  nextIpt = ref(""),
  hideIpt = ref("");
let focus = "";
const checkIn = (str: string, strs: string[]): boolean => {
  for (let i = 0; i < strs.length; i++) {
    const element = strs[i];
    if (element == str) {
      return true;
    }
  }
  return false;
};
onMounted(async () => {
  let res = await InitSetting();
  if (res.font_size) {
    for (const item of res.prev_group) {
      let has = Keys.find((s) => s.code == item);
      if (has) {
        prevPageKey.push(has.key);
        prevIpt.value = prevPageKey.join(" + ");
      }
    }
    for (const item of res.next_group) {
      let has = Keys.find((s) => s.code == item);
      if (has) {
        nextPageKey.push(has.key);
        nextIpt.value = nextPageKey.join(" + ");
      }
    }
    for (const item of res.hide_group) {
      let has = Keys.find((s) => s.code == item);
      if (has) {
        hideWindowsKey.push(has.key);
        hideIpt.value = hideWindowsKey.join(" + ");
      }
    }
    fontColor.value = res.font_color;
    fontSize.value = res.font_size;
    showSize.value = res.show_size;
  }
});
const clear = () => {
  prevPageKey = [];
  nextPageKey = [];
  hideWindowsKey = [];
  prevIpt.value = "";
  nextIpt.value = "";
  hideIpt.value = "";
  fontSize.value = 12;
  showSize.value = 50;
  fontColor.value = "rgb(220,220,220)";
};
document.onkeydown = function (event) {
  if (!focus) {
    return;
  }
  switch (focus) {
    case "prevPageKey":
      if (!checkIn(event.code, prevPageKey)) {
        prevPageKey.push(event.code);
        prevIpt.value = prevPageKey.join(" + ");
      }
      break;
    case "nextPageKey":
      if (!checkIn(event.code, nextPageKey)) {
        nextPageKey.push(event.code);
        nextIpt.value = nextPageKey.join(" + ");
      }
      break;
    case "hideWindowsKey":
      if (!checkIn(event.code, hideWindowsKey)) {
        hideWindowsKey.push(event.code);
        hideIpt.value = hideWindowsKey.join(" + ");
      }
      break;
  }
};
const prevPageFocus = () => {
  focus = "prevPageKey";
};
const nextPageFocus = () => {
  focus = "nextPageKey";
};
const hideWindowsFocus = () => {
  focus = "hideWindowsKey";
};
const loseBlur = () => {
  focus = "";
};

const saveSetting = async () => {
  let parms: any = {
    prev_group: [],
    next_group: [],
    hide_group: [],
    font_size: fontSize.value,
    font_color: fontColor.value,
    show_size: showSize.value,
  };
  for (const item of prevPageKey) {
    let exit = Keys.find((s) => s.key == item);
    if (exit) {
      parms.prev_group.push(exit.code);
    }
  }
  for (const item of nextPageKey) {
    let exit = Keys.find((s) => s.key == item);
    if (exit) {
      parms.next_group.push(exit.code);
    }
  }
  for (const item of hideWindowsKey) {
    let exit = Keys.find((s) => s.key == item);
    if (exit) {
      parms.hide_group.push(exit.code);
    }
  }
  let data = await SaveSetting(parms);
  if (data) {
    console.log("ok");
  } else {
    console.log("err");
  }
};
</script>

<style>

</style>