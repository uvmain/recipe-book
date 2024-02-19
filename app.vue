<script setup>
useHead({
  htmlAttrs: {
    lang: 'en',
  },
  charset: 'utf-8',
  title: 'RecipeBook',
})
</script>

<template>
  <div id="app" class="flex flex-col min-h-screen">
    <NuxtLoadingIndicator />
    <Header />
    <div class="flex-grow font-sans">
      <NuxtPage />
    </div>
    <Footer />
  </div>
</template>

<style>
html, body, #__nuxt {
  margin: 0;
  padding: 0;
  background: grey;
  color: white;
}
</style>
