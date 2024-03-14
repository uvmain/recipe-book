<template>
  <footer class="flex flex-col py-10 ml-8 text-zinc-700">
    <div class="text-xl my-2">
      RecipeBook
    </div>
    <div class="flex items-center gap-2">
      <div class="text-base">
        Made with
      </div>
      <a href="https://nuxt.com/" target="_blank" title="Nuxt 3">
        <NuxtLogo class="scale-50 -ml-6" />
      </a>
    </div>
    <div class="-m-8">
      <a href="https://github.com/uvmain/recipe-book" target="_blank" title="GitHub" class="flex items-center text-zinc-700 underline-none">
        <GithubLogo class="scale-30 bg-zinc-700 rounded-xl" />
        <div class="-ml-6 text-sm underline-none">
          https://github.com/uvmain/recipe-book
        </div>
      </a>
    </div>
  </footer>
</template>
