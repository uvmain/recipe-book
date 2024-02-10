<script setup lang="ts">
const siteTitle = 'RecipeBook'
</script>

<template>
  <header class="flex justify-between items-center p-4 bg-gray-100">
    <router-link to="/" class="no-underline">
      <h1 class="text-2xl font-bold text-gray px-10">
        {{ siteTitle }}
      </h1>
    </router-link>
  </header>
</template>
