export async function backendFetchRequest(path: string, options: RequestInit = {}): Promise<Response> {
  const url = `/api/${path}`
  const response = await fetch(url, options)
  return response
}
