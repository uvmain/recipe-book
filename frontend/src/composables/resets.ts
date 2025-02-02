import { useSessionStorage } from '@vueuse/core'

const searchInput = useSessionStorage<string>('searchInput', '')
const selectedCourses = useSessionStorage<string[]>('selectedCourses', [])
const selectedVegetarian = useSessionStorage<boolean>('selectedVegetarian', false)
const selectedCountry = useSessionStorage<string>('selectedCountry', '')
const selectedCalories = useSessionStorage<number>('selectedCalories', 1000)

export function resetSearch() {
  searchInput.value = ''
}

export function resetFilters() {
  selectedCourses.value = []
  selectedVegetarian.value = false
  selectedCalories.value = 1000
  selectedCountry.value = undefined
}
