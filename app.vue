<script setup>
useHead({
  htmlAttrs: {
    lang: 'en',
  },
  charset: 'utf-8',
  title: 'RecipeBook',
})

const isSidebarOpen = computed(() => {
  return useState('sidebarOpen').value
})
</script>

<template>
  <div id="app" class="bg-blue-gray-100 min-h-screen flex">
    <NuxtLoadingIndicator />
    <Sidebar :class="{'w-64': isSidebarOpen, 'w-0': !isSidebarOpen }"/>
    <div
      class="flex-1 antialiased"
      :class="{'pl-64 ml-4 md:ml-6 lg:ml-8': isSidebarOpen, 'pl-0': !isSidebarOpen }"
    >
      <NuxtPage />
    </div>
  </div>
</template>

<style>
html, body, #__nuxt {
  margin: 0;
  padding: 0;
  border: 0;
  font-family: 'Poppins', sans-serif;
}
</style>