<template>
  <footer class="flex flex-col py-10 pl-8 text-zinc-700">
    <div class="text-xl my-2">
      RecipeBook
    </div>
    <div class="flex flex-row items-center gap-2">
      <div class="text-base">
        Made with
      </div>
      <a href="https://nuxt.com/" target="_blank" title="Nuxt 3">
        <NuxtLogo class="h-6 opacity-70" />
      </a>
    </div>
    <div class="">
      <a href="https://github.com/uvmain/recipe-book" target="_blank" title="GitHub" class="flex flex-row items-center gap-2 text-zinc-700 underline-none">
        <GithubLogo class="w-8 h-8 bg-zinc-700 rounded-2xl p-1" />
        <div class="text-sm">
          https://github.com/uvmain/recipe-book
        </div>
      </a>
    </div>
  </footer>
</template>
