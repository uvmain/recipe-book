export async function getCachedImageUrl(imageFilename: string): Promise<string> {
  const url = imageFilename.length > 0 ? `/api/images/${imageFilename}` : '/default.webp'
  if ('caches' in window) {
    try {
      const cache = await caches.open('recipe-images')
      const cached = await cache.match(url)

      if (cached) {
        const blob = await cached.blob()
        return URL.createObjectURL(blob)
      }

      const response = await fetch(url)
      if (response.ok || response.type === 'opaque') {
        await cache.put(url, response.clone())
        const blob = await response.blob()
        return URL.createObjectURL(blob)
      }
    }
    catch (error) {
      console.error('Image caching failed:', error)
      return url
    }
  }
  return url
}
