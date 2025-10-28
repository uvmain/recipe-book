/// <reference types="vite-ssg" />

import { fileURLToPath, URL } from 'node:url'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Unfonts from 'unplugin-fonts/vite'
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'
import Components from 'unplugin-vue-components/vite'
import VueRouter from 'unplugin-vue-router/vite'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    VueRouter({
      /* options */
    }),
    AutoImport({
      imports: [
        'vue',
        'vue-router',
      ],
      dts: true,
      viteOptimizeDeps: true,
    }),
    vue(),
    Icons({
      scale: 1.0,
    }),
    // https://github.com/antfu/unplugin-vue-components
    Components({
      // allow auto load markdown components under `./src/components/`
      extensions: ['vue', 'md'],
      dts: true,
      include: [/\.vue$/, /\.vue\?vue/],
      // custom resolvers
      resolvers: [
        IconsResolver({
          prefix: 'icon',
        }),
      ],
    }),
    // https://github.com/antfu/unocss
    // see uno.config.ts for config
    UnoCSS(),
    Unfonts({
      google: {
        families: [
          'Onest',
        ],
      },
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  assetsInclude: [
    '**/*.md',
  ],
  ssgOptions: {
    script: 'async',
    formatting: 'minify',
  },
})
