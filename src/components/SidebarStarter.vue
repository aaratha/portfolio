<script setup lang="ts">
const emit = defineEmits<{
  (e: 'select', componentName: string): void
}>()

function selectContent(name: string) {
  emit('select', name)
}
</script>

<template>
  <div class="sidebar-start">
    <div class="sidebar">
      <button @click="selectContent('ContentA')">Show Content A</button>
      <button @click="selectContent('ContentB')">Show Content B</button>
    </div>
  </div>
</template>

<style scoped>
.sidebar-start {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
