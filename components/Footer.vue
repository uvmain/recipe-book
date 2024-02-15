<template>
  <footer class="flex flex-col p-10">
    <div class="text-xl my-2">
      RecipeBook
    </div>
    <div class="flex gap-2 items-center">
      <div class="opacity-50 text-sm">
        Made with
      </div>
      <a href="https://nuxt.com/" target="_blank" title="Nuxt 3">
        <NuxtLogo class="scale-50 -ml-6" />
      </a>
    </div>
    <div class="flex gap-2 items-center -m-6">
      <a href="https://github.com/uvmain/recipe-book" target="_blank" title="GitHub">
        <GithubLogo class="scale-20" />
      </a>
      <div class="-ml-6 opacity-50 text-sm">
        https://github.com/uvmain/recipe-book
      </div>
    </div>
  </footer>
</template>
