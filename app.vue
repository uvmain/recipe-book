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
  <div id="app" class="flex flex-col min-h-screen bg-primarybg-300">
    <NuxtLoadingIndicator />
    <Header />
    <div class="flex-grow font-sans antialiased">
      <NuxtPage />
    </div>
    <Footer />
  </div>
</template>

<style>
html, body, #__nuxt {
  margin: 0;
  padding: 0;
  color: white;
  border: 0;
}
</style>
