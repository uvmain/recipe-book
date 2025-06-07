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
    <Transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0 transform -translate-y-4"
      enter-to-class="opacity-100 transform translate-y-0"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="opacity-100 transform translate-y-0"
      leave-to-class="opacity-0 transform -translate-y-4"
    >
      <Filters v-if="showFiltersComponent" class="sticky z-10 background top-20" />
    </Transition>
    <RouterView />
  </div>
</template>

<style>
html, body, #app {
  margin: 0;
  padding: 0;
  border: 0;
  font-family: 'Onest', sans-serif;
  min-height: 100%;
}
</style>
