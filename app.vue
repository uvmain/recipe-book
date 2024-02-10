<template>
  <NuxtLoadingIndicator />
  <Header />
  <div class="h-full w-full font-sans">
    <NuxtPage />
  </div>
  <Footer />
</template>

<style>
html, body , #__nuxt{
  margin: 0;
  padding: 0;
  background: grey;
  color: white;
  color-scheme: dark;
}
</style>
