<script setup lang="ts">

function toggleSidebar() {
  isSidebarOpen.value = !isSidebarOpen.value
}
</script>

<template>
  <div class="fixed bottom-20 left-0 z-20">
    <button
      type="button"
      class="items-center h-12 text-center bg-white rounded-r-lg rounded-l-none shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-400"
      @click="toggleSidebar()"
    >
      <Icon v-if="isSidebarOpen" name="gg:chevron-double-left" class="size-6"/>
      <Icon v-else name="gg:chevron-double-right" class="size-6"/>
    </button>
  </div>
</template>