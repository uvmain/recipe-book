// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@unocss/nuxt',
    '@nuxt/content',
    'nuxt-icon',
    '@vueuse/nuxt',
  ],
  content: {},
  alias: {
    database: '/<srcDir>/public/database',
  },
})
