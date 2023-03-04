

<template>
  <div class="index">
    <n-grid cols="2">
      <n-grid-item>
        <div>
          <n-divider class="title font12" title-placement="left">
            选择书籍
          </n-divider>
          <n-row>
            <n-button type="info" dashed size="tiny" @click="selectFile">
              选择文本
            </n-button>
          </n-row>
          <br />
          <n-row>
            <n-scrollbar style="max-height: 220px">
              <n-list hoverable clickable>
                <n-list-item v-for="(item, i) in bookshelf">
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
                          {{ ~~(item.read_size / item.totle_size) }}%
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
let bookshelf: any = ref([]);
let props = defineProps<{
  mode: any;
}>();
const refBookshelf = async () => {
  let data = await GetBookshelf();
  bookshelf.value = data;
  console.log(data);
};
onMounted(async () => {
  refBookshelf();
});
const handleFinish = (res: any) => {
  refBookshelf();
};
const readBook = async (id: string) => {
  let content = await ReloadPage(id);
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
