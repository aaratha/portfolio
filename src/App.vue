<script setup lang="ts">
import SidebarMain from '@/components/SidebarMain.vue'
import SidebarStarter from '@/components/SidebarStarter.vue'
import ContentA from '@/components/ContentA.vue'
import ContentB from '@/components/ContentB.vue'
import { ref } from 'vue'

// null means no content selected, sidebarStart visible
const components = {
  ContentA,
  ContentB,
}

const selectedContent = ref<keyof typeof components | null>(null)

const sidebarTransitioning = ref(false)
const sidebarShown = ref(false)

function onSidebarAfterEnter() {
  sidebarTransitioning.value = false
  sidebarShown.value = true
}
function onSidebarBeforeEnter() {
  sidebarTransitioning.value = true
}
function onSidebarAfterLeave() {
  sidebarTransitioning.value = false
  sidebarShown.value = false
}
</script>

<template>
  <main>
    <!-- Sidebar transition -->
    <Transition
      name="fade"
      mode="out-in"
      @after-enter="onSidebarAfterEnter"
      @before-enter="onSidebarBeforeEnter"
      @after-leave="onSidebarAfterLeave"
    >
      <SidebarStarter v-if="!selectedContent" @select="selectedContent = $event" />
      <SidebarMain
        v-else
        :selected="selectedContent"
        @select="selectedContent = $event"
        @back="selectedContent = null"
      />
    </Transition>
    <!-- Content transition -->
    <Transition name="fade" mode="out-in">
      <component
        v-if="selectedContent && sidebarShown"
        :is="components[selectedContent]"
        :key="selectedContent"
        class="content"
      />
    </Transition>
  </main>
</template>
