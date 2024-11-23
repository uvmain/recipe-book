// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@unocss/nuxt',
    '@vueuse/nuxt',
    '@nuxt/eslint',
    "@nuxtjs/google-fonts",
    "@nuxt/icon"
  ],
  devServer: {
    port: 3001,
  },
  googleFonts: {
    families: {
      Poppins: true
    }
  },
  compatibilityDate: '2024-07-18'
})