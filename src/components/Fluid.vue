<template>
  <canvas id="canvas" ref="canvas" style="width: 100%; height: 100%" />
</template>

<script setup lang="ts">
import { onMounted, ref, nextTick } from 'vue'

// this will resolve to /public/pkg/wasm_fluid.js
import init from '@/wasm/wasm_fluid.js'

const canvas = ref<HTMLCanvasElement | null>(null)

onMounted(async () => {
  await nextTick() // ensure DOM is updated
  await init()
})
</script>

<style scoped>
canvas {
  position: absolute;
  top: 0;
  left: 0;
  z-index: -1;
}
</style>
