export function convertToAlphanumeric(inputString: string) {
  return inputString
    .replace(/[^a-z0-9]+/gi, '-') // Replace non-alphanumeric characters with a single hyphen
    .replace(/-+/g, '-') // Replace multiple hyphens with a single hyphen
    .toLowerCase() // Convert to lowercase
    .trim() // Trim leading and trailing spaces
}
