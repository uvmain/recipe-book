<script setup lang="ts">
import { countryToFlagCode } from '../composables/countries'

const props = defineProps({
  country: { type: String, required: false },
})

const flagCode = computed(() => {
  if (props.country === undefined || !(props.country in countryToFlagCode)) {
    return null
  }
  return countryToFlagCode[props.country].toUpperCase()
})
</script>

<template>
  <img v-if="flagCode" :src="`/flags/${flagCode}.svg`" />
</template>
