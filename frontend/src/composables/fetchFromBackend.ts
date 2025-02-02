export async function backendFetchRequest(path: string, options: RequestInit = {}): Promise<Response> {
  // const url = `/api/${path}`
  const url = `http://localhost:3001/api/${path}`
  const response = await fetch(url, options)
  return response
}
