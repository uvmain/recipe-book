import { defineConfig } from 'unocss'

export default defineConfig({
  shortcuts: {
    background: "bg-blue-gray-100 dark:bg-blue-gray-700",
    recipeCardBackground: "bg-white dark:bg-gray-600",
    border: "border-1 border-solid border-gray-200 dark:border-gray-500",
    text: "text-gray-800 dark:text-gray-200",
    recipeCardText: "text-gray-600 dark:text-gray-200",
    headerButton: "size-12 bg-white dark:bg-gray-600 overflow-hidden shadow-md hover:shadow-lg dark:hover:shadow-gray-600 transition-shadow duration-100 active:bg-blue-gray-300 border rounded",
    headerSearch: "h-11 text-xl focus:placeholder-op-0 focus:bg-blue-gray-300 text-center md:text-2xl px-2  outline-none w-1/2 md:w-1/4 dark:placeholder-gray-300",
    headerButtonIcon: "text-2xl align-middle text-center text",
    titleText: "text-dark dark:text-white",
    filterHr: "w-80% my-2 border-gray-300 dark:border-gray-500 border-1 border-solid"
  },
  theme: {
    colors: {},
  },
})
