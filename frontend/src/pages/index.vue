

<template>
  <div class="index"  style="--wails-draggable: drag">
    <n-grid cols="2">
      <n-grid-item>
        <div>
          <n-divider class="title font12" title-placement="left">
            选择书籍
          </n-divider>
          <n-row>
            <n-upload action="/api/uploadFile">
              <n-button>选择文件</n-button>
            </n-upload>
            文本需要设置为UTF-8编码
          </n-row>
          <n-row style="text-align: center">
            <n-button strong secondary type="info" size="small" @click="ReadMod">
              阅读模式
            </n-button>
          </n-row>
        </div>
      </n-grid-item>
      <n-grid-item> <Setting></Setting> </n-grid-item>
    </n-grid>
  </div>
</template>
<script lang="ts" setup>
import Setting from "./setting.vue";
import { ReadMod } from "../../wailsjs/go/read/App";

</script>
<style>
.index {
  padding: 5px 12px;
}
.key {
  line-height: 28px;
}
</style>
