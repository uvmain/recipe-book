<script setup lang="ts">
const vulgarFractions = [
  '½',
  '⅓',
  '¼',
  '⅕',
  '⅙',
  '⅐',
  '⅛',
  '⅑',
  '⅒',
  '⅔',
  '⅖',
  '¾',
  '⅗',
  '⅘',
  '⅚',
  '⅜',
  '⅝',
  '⅞',
]
</script>

<template>
  <div class="grid gap-2 grid-cols-5">
    <div v-for="(vulgarFraction, index) in vulgarFractions" :key="index" class="text-lg recipeCardBackground text-center rounded">
      {{ vulgarFraction }}
    </div>
  </div>
</template>
