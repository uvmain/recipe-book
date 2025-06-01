/// <reference types="vite-ssg" />

import { fileURLToPath, URL } from 'node:url'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Unfonts from 'unplugin-fonts/vite'
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'
import Components from 'unplugin-vue-components/vite'
import { defineConfig } from 'vite'
import Pages from 'vite-plugin-pages'
import { VitePWA } from 'vite-plugin-pwa'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    AutoImport({
      imports: [
        'vue',
        'vue-router',
      ],
      dts: true,
      viteOptimizeDeps: true,
    }),
    vue(),
    VitePWA(
      {
        manifest: {
          name: 'RecipeBook',
          short_name: 'RecipeBook',
          icons: [
            {
              src: '/favicon-256.png',
              sizes: '256x256',
              type: 'image/png',
              purpose: 'any maskable',
            },
          ],
        },
        includeAssets: [
          '/audio/sonnette_reveil.wav',
          '/default.webp',
          'favicon.ico',
        ],
        workbox: {
          runtimeCaching: [
            {
              urlPattern: ({ url }) => {
                return url.pathname.startsWith('/api') && !['/api/check-session', '/api/random-recipe'].includes(url.pathname)
              },
              handler: 'CacheFirst' as const,
              options: {
                cacheName: 'api-cache',
                cacheableResponse: {
                  statuses: [0, 200, 206],
                },
              },
            },
            {
              urlPattern: ({ url }) => {
                return url.pathname.startsWith('/api/images')
              },
              handler: 'CacheFirst' as const,
              options: {
                cacheName: 'image-cache',
                expiration: {
                  maxEntries: 300,
                  maxAgeSeconds: 30 * 24 * 60 * 60,
                },
                cacheableResponse: {
                  statuses: [0, 200],
                },
              },
            },
          ],
        },
      },
    ),
    // https://github.com/hannoeru/vite-plugin-pages
    Pages(),
    Icons({
      scale: 1.0,
    }),
    // https://github.com/antfu/unplugin-vue-components
    Components({
      // allow auto load markdown components under `./src/components/`
      extensions: ['vue', 'md'],
      dts: true,
      // allow auto import and register components used in markdown
      include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
      // custom resolvers
      resolvers: [
        // auto import icons
        // https://github.com/antfu/unplugin-icons
        IconsResolver({
          prefix: 'icon',
          // enabledCollections: ['carbon']
        }),
      ],
    }),
    // https://github.com/antfu/unocss
    // see uno.config.ts for config
    UnoCSS(),
    Unfonts({
      google: {
        families: [
          'Quicksand',
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
    format: 'cjs',
    formatting: 'minify',
    beastiesOptions: {
      reduceInlineStyles: false,
    },
  },
})
