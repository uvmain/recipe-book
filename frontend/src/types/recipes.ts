export interface Recipe {
  slug: string
  date_created: string
  name: string
  author: string
  source: string
  course: string
  country: string
  vegetarian: boolean
  prep_time: string
  cooking_time: string
  calories: string
  servings: string
  ingredients: string
  instructions: string
  imageFilename: string
  imageWidth: number
  imageHeight: number
}

export interface RecipeCard {
  slug: string
  date_created: string
  name: string
  author: string
  source: string
  course: string
  country: string
  vegetarian: boolean
  prep_time: string
  cooking_time: string
  calories: string
  servings: string
  imageFilename: string
  imageWidth: number
  imageHeight: number
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
