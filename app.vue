<script setup>
import { isSidebarOpen } from './utils/sidebar'

useHead({
  htmlAttrs: {
    lang: 'en',
  },
  charset: 'utf-8',
  title: 'RecipeBook',
})

const leftMargin = computed(() => {
  return isSidebarOpen.value
})

</script>

<template>
  <div id="app" class="bg-blue-gray-100 min-h-screen flex">
    <NuxtLoadingIndicator />
    <Sidebar />
    <SidebarToggle />
    <NuxtPage :class="{ 'ml-68 md:ml-70 lg:ml-72' : leftMargin }" />
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