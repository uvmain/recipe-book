<script setup>
useHead({
  htmlAttrs: {
    lang: 'en',
  },
  charset: 'utf-8',
  title: 'RecipeBook',
})
</script>

<template class="">
  <div id="app" class="bg-gray-400 min-h-screen">
    <NuxtLoadingIndicator />
    <Header />
    <div class="antialiased">
      <NuxtPage />
    </div>
    <Footer class=" md:bottom-0 md:left-0"/>
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
