export async function base64ToBlob(base64: string): Promise<Blob> {
  const response = await fetch(base64)
  const newBlob = await response.blob()
  return newBlob
}
