/// <reference lib="WebWorker" />
/// <reference types="vite/client" />
import type { PrecacheEntry } from 'workbox-precaching'
import { CacheableResponsePlugin } from 'workbox-cacheable-response'
import { ExpirationPlugin } from 'workbox-expiration'
import { cleanupOutdatedCaches, createHandlerBoundToURL, precacheAndRoute } from 'workbox-precaching'
import { NavigationRoute, registerRoute } from 'workbox-routing'
import { StaleWhileRevalidate } from 'workbox-strategies'

declare const self: ServiceWorkerGlobalScope

interface InvalidateCacheMessage {
  type: 'INVALIDATE_CACHE'
  url: string
}

self.addEventListener('message', (event) => {
  const data = event.data as InvalidateCacheMessage

  if (data?.type === 'INVALIDATE_CACHE' && data.url) {
    const url = new URL(data.url, self.location.origin)

    void caches.open('recipebook-cache')
      .then(async cache => cache.delete(url))
      .then((success) => {
        if (success) {
          console.log(`[SW] Cache invalidated for: ${url}`)
        }
        else {
          console.warn(`[SW] No cache entry found for: ${url}`)
        }
      })
      .catch((err) => {
        console.error(`[SW] Failed to invalidate cache for: ${url}`, err)
      })
  }
})

const entries = self.__WB_MANIFEST as PrecacheEntry[]

if (!entries.some(entry => entry.url === '/')) {
  entries.push({ url: '/', revision: Math.random().toString() })
}

precacheAndRoute(entries)

cleanupOutdatedCaches()

registerRoute(
  ({ sameOrigin, url }) =>
    sameOrigin
    && !(url.pathname === '/api/check-session')
    && !(url.pathname === '/api/login')
    && !(url.pathname === '/api/logout')
    && !(url.pathname === '/api/random-recipe'),
  new StaleWhileRevalidate({
    cacheName: 'recipebook-cache',
    plugins: [
      new CacheableResponsePlugin({ statuses: [200] }),
      // 2 months max
      new ExpirationPlugin({ purgeOnQuotaError: true, maxAgeSeconds: 60 * 60 * 24 * 60 }),
    ],
  }),
)

// to allow work offline
registerRoute(new NavigationRoute(
  createHandlerBoundToURL(import.meta.env.BASE_URL || '/'),
))
