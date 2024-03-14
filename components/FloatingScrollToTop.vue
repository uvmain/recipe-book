<script setup lang="ts">
function scrollToTop() {
  window.scrollTo(0, 0)
}
</script>

<template>
  <button
    type="button"
    class="text-white text-center font-medium md:text-3xl text-base bg-blue-gray-400 fixed bottom-2 md:bottom-10 md:right-10 z-10 right-2 rounded-xl border-0"
    @click="scrollToTop"
  >
    <Icon name="carbon:arrow-up" />
  </button>
</template>
