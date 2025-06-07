import { defineConfig } from 'unocss'

export default defineConfig({
  shortcuts: {
    background: 'bg-slate-100 dark:bg-slate-900',
    recipeCardBackground: 'bg-white dark:bg-slate-800',
    border: 'border-1 border-solid border-slate-300 dark:border-slate-600',
    text: 'text-slate-700 dark:text-slate-200',
    recipeCardText: 'text-slate-600 dark:text-slate-300',
    headerButton: 'size-12 bg-white dark:bg-slate-700 overflow-hidden border border-slate-300 dark:border-slate-600 rounded hover:cursor-pointer text-slate-700 dark:text-slate-200 transition-colors transition-transform duration-150 ease-in-out hover:bg-slate-200 dark:hover:bg-slate-600 active:scale-95 active:bg-slate-300 dark:active:bg-slate-500',
    textButton: 'size-12 bg-white dark:bg-slate-700 overflow-hidden border border-slate-300 dark:border-slate-600 rounded text-slate-700 dark:text-slate-200 transition-colors transition-transform duration-150 ease-in-out hover:bg-slate-200 dark:hover:bg-slate-600 active:scale-95 active:bg-slate-300 dark:active:bg-slate-500',
    headerSearch: 'h-11 text-xl focus:placeholder-op-0 focus:bg-slate-300 dark:focus:bg-slate-600 bg-white dark:bg-slate-700 text-slate-700 dark:text-slate-200 text-center md:text-2xl px-2 outline-none w-1/2 md:w-1/4 placeholder-slate-400 dark:placeholder-slate-400',
    headerButtonIcon: 'text-2xl align-middle text-center text-slate-700 dark:text-slate-200',
    titleText: 'text-slate-800 dark:text-slate-100',
    filterHr: 'w-80% my-4 border-slate-300 dark:border-slate-600 border-1 border-solid',
  },
  theme: {
    colors: {},
  },
})
