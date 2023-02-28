<script lang="ts" setup>
import { ref, onMounted, reactive } from "vue";
import { WindowMinimise } from "../../wailsjs/runtime/runtime";
const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
});
</script>

<template>
  <main >
    <div style="--wails-dragable: drag;height: 100px;width: 50px;background-color: aqua;"></div>
    <div id="result" class="result">{{ data.resultText }}</div>
    <button @click="WindowMinimise">最小化</button>
    <n-upload
      action="/api/uploadFile"
      :headers="{
        'naive-info': 'hello!',
      }"
      :data="{
        'naive-data': 'cool! naive!',
      }"
    >
      <n-button>选择</n-button>
    </n-upload>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
}
</style>
