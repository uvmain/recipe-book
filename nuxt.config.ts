// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@unocss/nuxt',
    'nuxt-icon',
    '@vueuse/nuxt',
    '@nuxt/eslint'
  ],
  devServer: {
    port: 3001,
  },
})