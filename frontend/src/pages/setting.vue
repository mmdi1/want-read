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
            <n-button strong secondary type="info" size="small">
              保存
            </n-button>
          </n-space>
        </n-col>
      </n-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue";
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
const clear = () => {
  prevPageKey = [];
  nextPageKey = [];
  hideWindowsKey = [];
  prevIpt.value = "";
  nextIpt.value = "";
  hideIpt.value = "";
};
document.onkeydown = function (event) {
  if (!focus) {
    return;
  }
  switch (focus) {
    case "prevPageKey":
      if (!checkIn(event.key, prevPageKey)) {
        prevPageKey.push(event.key);
        prevIpt.value = prevPageKey.join(" + ");
      }
      break;
    case "nextPageKey":
      if (!checkIn(event.key, nextPageKey)) {
        nextPageKey.push(event.key);
        nextIpt.value = nextPageKey.join(" + ");
      }
      break;
    case "hideWindowsKey":
      if (!checkIn(event.key, hideWindowsKey)) {
        hideWindowsKey.push(event.key);
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
</script>

<style></style>