import { defineConfig } from 'unocss'

export default defineConfig({
  shortcuts: {
    background: 'bg-stone-50 dark:bg-zinc-800',
    recipeCardBackground: 'bg-white dark:bg-stone-700',
    border: 'border-1 border-solid border-stone-200 dark:border-stone-500',
    text: 'text-stone-800 dark:text-stone-200',
    recipeCardText: 'text-stone-600 dark:text-stone-200',
    headerButton: 'size-12 bg-white dark:bg-stone-600 overflow-hidden shadow-md hover:shadow-lg dark:shadow-stone-700 dark:hover:shadow-stone-600 transition-shadow duration-100 active:bg-stone-300 border rounded',
    textButton: 'size-12 bg-white dark:bg-stone-600 overflow-hidden shadow-md hover:shadow-lg dark:shadow-stone-700 dark:hover:shadow-stone-600 transition-shadow duration-100 active:bg-stone-300 border rounded text',
    headerSearch: 'h-11 text-xl focus:placeholder-op-0 focus:bg-stone-300 text-center md:text-2xl px-2  outline-none w-1/2 md:w-1/4 dark:placeholder-stone-200',
    headerButtonIcon: 'text-2xl align-middle text-center text',
    titleText: 'text-dark dark:text-white',
    filterHr: 'w-80% my-4 border-stone-300 dark:border-stone-500 border-1 border-solid',
  },
  theme: {
    colors: {},
  },
})
