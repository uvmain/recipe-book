<script setup>
const input = ref('')

function filteredList() {
  return ['apple', 'banana', 'orange'].filter(fruit =>
    fruit.toLowerCase().includes(input.value.toLowerCase()),
  )
}

const { data: recipes } = await useAsyncData('allrecipes', () =>
  queryContent('/recipes').find())

function flattenRecipe(jsonObj) {
  const result = []
  function traverse(obj) {
    for (const key in obj) {
      if (typeof obj[key] === 'object' && obj[key] !== null) {
        traverse(obj[key])
      }
      else if (key === 'vegetarian' && obj[key] === true) {
        result.push('vegetarian')
      }
      else if (typeof obj[key] === 'string') {
        result.push(obj[key].replace(/\W/g, ' '))
      }
    }
  }
  traverse(jsonObj)
  return result
}

function onlyUnique(value, index, array) {
  return array.indexOf(value) === index
}

const slugKeywords = computed(() => {
  const slugKeywords = []
  for (const recipe of recipes.value) {
    const flatRecipe = flattenRecipe(recipe)
    console.log(flatRecipe)
    slugKeywords.push({
      [recipe.slug]: flatRecipe.join(' ').split(' ').filter(word => word.length > 2).filter(onlyUnique),
    })
  }
  return slugKeywords
})
</script>

<template>
  <div v-if="recipes" class="text-black">
    {{ Object.values(slugKeywords[0]) }}
  </div>
  <input v-model="input" type="text" placeholder="Search fruits...">
  <div v-if="input && !filteredList().length" class="item error">
    <p>No results found!</p>
  </div>
</template>
