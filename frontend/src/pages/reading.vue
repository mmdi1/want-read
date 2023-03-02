<template>
  <div id="reading">
    {{ props.mode.content }}
  </div>
</template>

<script lang="ts" setup>
let props = defineProps<{
  mode: any;
}>();
</script>

<style>
.reading {
  word-break: break-all;
}
</style>