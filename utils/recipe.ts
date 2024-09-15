export interface Recipe {
  slug: string
  date_created: string
  name: string
  author: string
  source: string
  course: string
  vegetarian: boolean
  prep_time: string
  cooking_time: string
  calories: string
  servings: string
  ingredients: string
  instructions: string
  image: string
  country: string
}

export const allowedCourses = [
  'baking',
  'breakfasts',
  'cocktails',
  'desserts',
  'mains',
  'salads',
  'sauces',
  'sides',
  'soups',
  'snacks',
  'seasonings',
]

export interface RecipesApiResponse {
  data: Recipe[]
}