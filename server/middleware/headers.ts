export default defineEventHandler((event) => {
  const url = event.node.req.url

  if (url?.startsWith('/api/images/') || url?.startsWith('/api/thumbnail/')) {
    setResponseHeaders(event, {
      'Cache-Control': 'public, max-age=31536000, immutable',
      Expires: new Date(Date.now() + 31536000000).toUTCString(),
    })
  }
})