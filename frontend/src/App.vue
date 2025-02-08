<script setup>
import { useSessionStorage } from '@vueuse/core'
import { useHead } from '@vueuse/head'

const route = useRoute()

const showFilters = useSessionStorage('showFilters', false)

const showFiltersComponent = computed(() => {
  return route.path === '/' && showFilters.value
})

useHead({
  htmlAttrs: {
    lang: 'en',
  },
  charset: 'utf-8',
  title: 'RecipeBook',
  meta: [
    { name: 'description', content: 'Personal Recipe Book' },
  ],
})
</script>

<template>
  <div id="app" class="min-h-screen flex flex-col background">
    <Header class="sticky top-0 z-10" />
    <Filters v-if="showFiltersComponent" class="sticky z-10 background top-20" />
    <RouterView />
  </div>
</template>

<style>
html, body, #app {
  margin: 0;
  padding: 0;
  border: 0;
  font-family: 'Quicksand', sans-serif;
  min-height: 100%;
  @apply standard;
}
</style>
