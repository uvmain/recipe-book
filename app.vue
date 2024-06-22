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
  <div id="app" class="bg-primary_bg-300 min-h-screen">
    <NuxtLoadingIndicator />
    <Header />
    <div class="font-sans antialiased">
      <NuxtPage />
    </div>
    <Footer />
  </div>
</template>

<style>
html, body, #__nuxt {
  margin: 0;
  padding: 0;
  background: rgb(209, 213, 219);
  color: white;
  border: 0;
}
</style>
