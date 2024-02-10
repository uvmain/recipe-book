<template>
  <nav class="bg-gray-800 text-white p-4 flex justify-between items-center">
    <div class="text-2xl font-bold mb-4">
      <router-link to="/" class="text-white">
        Recipes
      </router-link>
    </div>

    <div>
      <router-link to="/" class="text-white mr-1rem transition hover:text-blue-300">
        Home
      </router-link>
      <router-link to="/recipes" class="nav-link">
        Recipes
      </router-link>
    </div>
  </nav>
</template>
