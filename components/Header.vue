<script setup lang="ts">
const siteTitle = 'RecipeBook'
</script>

<template>
  <header class="flex justify-between items-center p-4 bg-gray-100">
    <div>
      <h1 class="text-2xl font-bold text-gray px-10">
        {{ siteTitle }}
      </h1>
    </div>
    <nav class="main-nav flex gap-4">
      <router-link to="/" class="text-gray-600">
        Home
      </router-link>
      <router-link to="/" class="text-gray-600">
        Recipes
      </router-link>
    </nav>
  </header>
</template>
