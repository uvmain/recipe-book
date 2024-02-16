export interface Recipe {
  slug: string
  dateAdded: string
  name: string
  author: string
  source: string
  course: string
  country: string
  vegetarian: false
  prepTime: string
  cookingTime: string
  calories: string
  servings: string
  ingredients: string[]
  instructions: string[]
  image: string
}

export const allowedCourses = [
  'baking',
  'breakfasts',
  'mains',
  'desserts',
  'sauces',
  'sides',
  'soups',
  'snacks',
  'cocktails',
]
