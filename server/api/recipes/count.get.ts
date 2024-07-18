const baseUrl = process.env.DIRECTUS_URL as string;
const token = process.env.DIRECTUS_TOKEN as string;

export default defineEventHandler(async () => {
  if (baseUrl.length < 1 || token.length < 1) {
    return {
      error: 'CMS connection variables are not defined'
    };
  }
  
  const url = `${baseUrl}/items/Recipe?aggregate[count]=*`;

  try {
    const options = {
      headers: {
        Authorization: `Bearer ${token}`
      },
    };
    const response = await $fetch<{ data:[{ count: number }]}>(url, options).catch((error) => {
      throw new Error(`Failed to fetch data: ${JSON.stringify(error.data)}`);
    });

    return response.data[0];
  }
 catch (error) {
    return { error: `Failed to fetch data from CMS: ${error}` };
  }
});
