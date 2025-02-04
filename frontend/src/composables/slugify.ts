export function slugify(stringToSlug: string) {
  stringToSlug = stringToSlug.replace(/^\s+|\s+$/g, '') // trim leading/trailing white space
  stringToSlug = stringToSlug.toLowerCase() // convert string to lowercase
  stringToSlug = stringToSlug.replace(/[^a-z0-9 -]/g, '') // remove any non-alphanumeric characters
    .replace(/\s+/g, '-') // replace spaces with hyphens
    .replace(/-+/g, '-') // remove consecutive hyphens
  return stringToSlug
}
