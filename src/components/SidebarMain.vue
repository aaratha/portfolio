<script setup lang="ts">
const props = defineProps<{
  selected: string
}>()

const emit = defineEmits<{
  (e: 'select', componentName: string): void
  (e: 'back'): void
}>()

function selectContent(name: string) {
  if (name !== props.selected) emit('select', name)
}

function goBack() {
  emit('back')
}
</script>

<template>
  <div class="sidebar fade-in">
    <button @click="goBack">Back</button>
    <button @click="selectContent('ContentA')">Content A</button>
    <button @click="selectContent('ContentB')">Content B</button>
  </div>
</template>

<style scoped></style>
