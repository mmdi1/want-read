<script lang="ts" setup>
import { onMounted, reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { WindowMinimise } from '../../wailsjs/runtime/runtime'
const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

const handleFolderSelect = (event: any) => {
  var files = event.target.files;
  console.log("aaaaaaaaaaaaaaaa")
  if (files.length > 0) {
    var folderPath = files[0].path;
    console.log(folderPath);
  }
}
onMounted(() => {
  let ele = document.getElementById("open")
  console.log('aaaaaaaaaaaaaaaaa', ele)
  ele?.addEventListener('change', handleFolderSelect);
})
</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <button @click="WindowMinimise">最小化</button>
    <input type="file" name="filename" id="open" />
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
}
</style>
