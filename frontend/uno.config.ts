import { defineConfig } from 'unocss'

export default defineConfig({
  shortcuts: {
    background: 'bg-zinc-100 dark:bg-zinc-800',
    recipeCardBackground: 'bg-white dark:bg-zinc-700',
    border: 'border-1 border-solid border-zinc-300 dark:border-zinc-600',
    text: 'text-zinc-700 dark:text-zinc-200',
    recipeCardText: 'text-zinc-600 dark:text-zinc-300',
    headerButton: 'size-12 bg-white dark:bg-zinc-700 overflow-hidden border border-zinc-300 dark:border-zinc-600 rounded hover:cursor-pointer text-zinc-700 dark:text-zinc-200 transition-colors transition-transform duration-150 ease-in-out hover:bg-zinc-200 dark:hover:bg-zinc-600 active:scale-95 active:bg-zinc-300 dark:active:bg-zinc-500',
    textButton: 'size-12 bg-white dark:bg-zinc-700 overflow-hidden border border-zinc-300 dark:border-zinc-600 rounded text-zinc-700 dark:text-zinc-200 transition-colors transition-transform duration-150 ease-in-out hover:bg-zinc-200 dark:hover:bg-zinc-600 active:scale-95 active:bg-zinc-300 dark:active:bg-zinc-500',
    headerSearch: 'h-11 text-xl focus:placeholder-op-0 focus:bg-zinc-300 dark:focus:bg-zinc-600 bg-white dark:bg-zinc-700 text-zinc-700 dark:text-zinc-200 text-center md:text-2xl px-2 outline-none w-1/2 md:w-1/4 placeholder-zinc-400 dark:placeholder-zinc-400',
    headerButtonIcon: 'text-2xl align-middle text-center text-zinc-700 dark:text-zinc-200',
    titleText: 'text-zinc-800 dark:text-zinc-100',
    filterHr: 'w-80% my-4 border-zinc-300 dark:border-zinc-600 border-1 border-solid',
  },
  theme: {
    colors: {},
  },
})
