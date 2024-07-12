<script setup lang="ts">
function scrollToTop() {
  window.scrollTo(0, 0)
}
</script>

<template>
  <button
    type="button"
    class="text-base text-white text-center font-medium md:text-3xl bg-blue-gray-400 fixed bottom-2 md:bottom-10 md:right-10 z-10 right-2 rounded-xl border-1 border-solid border-white hover:bg-blue-gray-600 hover:shadow-lg"
    @click="scrollToTop"
  >
    <Icon name="carbon:arrow-up" />
  </button>
</template>
