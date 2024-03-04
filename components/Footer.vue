<template>
  <footer class="flex flex-col p-10">
    <div class="text-xl my-2">
      RecipeBook
    </div>
    <div class="flex gap-2 items-center">
      <div class="text-sm opacity-50">
        Made with
      </div>
      <a href="https://nuxt.com/" target="_blank" title="Nuxt 3">
        <NuxtLogo class="scale-50 -ml-6" />
      </a>
    </div>
    <div class="-m-6">
      <a href="https://github.com/uvmain/recipe-book" target="_blank" title="GitHub" class="flex gap-2 items-center text-white">
        <GithubLogo class="scale-20" />
        <div class="-ml-6 opacity-50 text-sm">
          https://github.com/uvmain/recipe-book
        </div>
      </a>
    </div>
  </footer>
</template>
