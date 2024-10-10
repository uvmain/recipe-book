export function resetSearch() {
  useState<string>('searchInput').value = ''
}

export function resetFilters() {
  useState<string[]>('selectedCourses').value = []
  useState<boolean>('selectedVegetarian').value = false
  useState<number>('selectedCalories').value = 1000
  useState('selectedCountry').value = undefined
}