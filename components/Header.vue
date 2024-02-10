<script setup lang="ts">
const siteTitle = 'RecipeBook'
</script>

<template>
  <header class="flex justify-between items-center p-4 bg-gray-100">
    <div>
      <h1 class="text-2xl font-bold text-gray">
        {{ siteTitle }}
      </h1>
    </div>
    <nav class="main-nav flex gap-4">
      <router-link to="/" class="text-blue-500">
        Home
      </router-link>
      <router-link to="/recipes" class="text-blue-500">
        Recipes
      </router-link>
    </nav>
  </header>
</template>
