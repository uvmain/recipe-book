const baseUrl = process.env.DIRECTUS_URL as string;
const token = process.env.DIRECTUS_TOKEN as string;
const cache: { [key: string]: any } = {};

export default defineEventHandler(async (event) => {
  const slug = event.context.params?.slug;

  if (!slug) {
    return { error: 'Slug is undefined' };
  }

  const cacheKey = Array.isArray(slug) ? slug.join('/') : slug;
  const url = `${baseUrl}/assets/${cacheKey}?width=800&height=800&fit=inside`;

  if (cache[cacheKey]) {
    return cache[cacheKey];
  }

  try {
    const response = await $fetch(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }).catch((error) => {
      throw new Error(`Failed to fetch data: ${JSON.stringify(error.data)}`);
    });

    cache[cacheKey] = response;

    return response;
  }
 catch (error) {
    return { error: `Failed to fetch data from CMS: ${error}` };
  }
});
