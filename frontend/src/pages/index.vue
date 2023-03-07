

<template>
  <div class="index">
    <n-grid cols="2">
      <n-grid-item>
        <div>
          <n-divider class="title font12" title-placement="left">
            选择书籍
          </n-divider>
          <n-row>
            <n-space justify="space-around">
              <n-button type="info" dashed size="tiny" @click="selectFile">
                选择文本
              </n-button>
              <n-tooltip
                :show-arrow="false"
                placement="bottom-start"
                trigger="hover"
              >
                <template #trigger>
                  <n-tag size="small">使用说明</n-tag>
                </template>
                <div style="width: 300px; font-size: 12px; text-align: left">
                  1.选择txt书籍上传。
                  <br />
                  2.点击阅读即可开始阅读，首次阅读模式，会有很大部分隐藏框，
                  <span style="color: yellowgreen"
                    >将鼠标移动至第一行字体头顶调整边框大小</span
                  >，二次打开将自动记忆。
                  <br />
                  3.当前字数后输入框是已读的字数，手动输入后点击阅读即可直接跳转字数位置。
                  <br />
                  4.设置界面可自行配置习惯用键，保存则生效，恢复初始则点击初始后保存即可。
                  <br />
                  5.阅读模式下，双击字体即可唤出主界面，或右下角小托盘点击设置即可。
                  <br />
                  6.程序是单机应用，无联网操作，放心使用。
                </div>
              </n-tooltip>
            </n-space>
          </n-row>
          <br />
          <n-row>
            <n-scrollbar style="max-height: 220px">
              <n-list hoverable clickable>
                <n-list-item v-for="(item, i) in props.mode.bookshelf">
                  <n-thing content-style="margin-top: 10px;">
                    <template #description>
                      <p class="font12 title">{{ item.name }}</p>
                      <n-space size="small" style="margin-top: 4px">
                        <n-button
                          type="warning"
                          size="tiny"
                          @click="readBook(item.id)"
                        >
                          阅读
                        </n-button>
                        <n-tag :bordered="false" type="info" size="small">
                          {{ ((item.read_size?((item.read_size) / (item.total_size))*100:0) as any).toFixed(2)



                          }}%
                        </n-tag>
                        <n-tag :bordered="false" type="info" size="small">
                          {{ item.total_size }}字
                        </n-tag>
                        <n-button
                          type="error"
                          size="tiny"
                          @click="removeBook(item.id)"
                        >
                          删除
                        </n-button>
                        <br />
                        <n-tag :bordered="false" type="info" size="small">
                          当前字数
                        </n-tag>
                        <n-input-number
                          v-model:value="item.read_size"
                          :show-button="false"
                          size="tiny"
                        />
                      </n-space>
                    </template>
                  </n-thing>
                </n-list-item>
              </n-list>
            </n-scrollbar>
          </n-row>
        </div>
      </n-grid-item>
      <n-grid-item> <Setting></Setting> </n-grid-item>
    </n-grid>
  </div>
</template>
<script lang="ts" setup>
import Setting from "./setting.vue";
import {
  GetBookshelf,
  ReloadPage,
  RemoveBook,
  SelectFile,
} from "../../wailsjs/go/read/App";
import { onMounted, ref } from "vue";
let props = defineProps<{
  mode: any;
}>();
const refBookshelf = async () => {
  let data = await GetBookshelf();
  props.mode.bookshelf = data;
};
onMounted(async () => {
  refBookshelf();
});
const readBook = async (id: string) => {
  let readBook = props.mode.bookshelf.find((s: any) => s.id == id);
  let content = await ReloadPage(id, readBook.read_size);
  props.mode.screen = "read";
  props.mode.content = content;
};
const removeBook = async (id: string) => {
  let ok = await RemoveBook(id);
  if (ok) {
    refBookshelf();
  }
};
const selectFile = async () => {
  let str = await SelectFile();
  if (str) {
    refBookshelf();
  }
};
</script>
<style>
.index {
  padding: 5px 12px;
}
.key {
  line-height: 28px;
}
.n-row {
  display: block;
}
.title {
  text-align: left;
  margin-block-start: 0;
  margin-block-end: 0;
}
</style>
