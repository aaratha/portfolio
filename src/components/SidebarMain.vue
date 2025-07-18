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

// function goBack() {
//   emit('back')
// }
</script>

<template>
  <div class="sidebar fade-in">
    <!-- <button @click="goBack">Back</button> -->
    <button class="sidebar-button" @click="selectContent('Fluid')">Projects</button>
    <div class="sub-sidebar">
      <button class="sub-sidebar-button" @click="selectContent('Fluid')">Fluid Simulation</button>
      <button class="sub-sidebar-button" @click="selectContent('PinIt')">PinIt!</button>
      <button class="sub-sidebar-button" @click="selectContent('Strings')">Strings</button>
    </div>
    <button class="sidebar-button" @click="selectContent('Climate')">Papers</button>
    <div class="sub-sidebar">
      <button class="sub-sidebar-button" @click="selectContent('Climate')">
        Anthropogenic Climate Change
      </button>
      <button class="sub-sidebar-button" @click="selectContent('Fractal')">
        Fractal Viscous Fingering
      </button>
      <button class="sub-sidebar-button" @click="selectContent('Superconductor')">
        High Temp Superconductor
      </button>
    </div>
    <button class="sidebar-button" @click="selectContent('Photos')">Artworks</button>
    <div class="sub-sidebar">
      <button class="sub-sidebar-button" @click="selectContent('Photos')">Photos</button>
      <button class="sub-sidebar-button" @click="selectContent('Illustrations')">
        Illustrations
      </button>
      <button class="sub-sidebar-button" @click="selectContent('Films')">Films</button>
    </div>
  </div>
</template>

<style scoped></style>
