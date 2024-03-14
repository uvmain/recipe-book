import { defineConfig } from 'unocss'

export default defineConfig({
  shortcuts: {
    'add-form-component': 'block w-full p-3 text-base font-normal text-dark bg-white bg-clip-padding border border-solid border-primarybg-300 rounded transition ease-in-out m-0',
    'tip-tap-button': 'block px-2 py-1 text-dark text-base bg-white bg-clip-padding border border-solid border-primarybg-300 rounded-md transition ease-in-out m-0',
    'header-button': 'text-white text-xl hover:bg-zinc-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 border-1 border-solid border-zinc-700 bg-blue-gray-400',
    'header-button-selected': 'text-white text-xl bg-zinc-400 bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 border-1 border-solid border-zinc-700',
    'search-bar': 'text-left text-zinc-600 text-xl bg-blue-gray-400 hover:bg-zinc-300 focus:bg-zinc-200 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 border-1 border-solid border-zinc-700 outline-none w-1/2 md:w-1/5',
  },
  theme: {
    colors: {
      primarybg: {
        300: '#d1d5db',
      },
    },
  },
})
