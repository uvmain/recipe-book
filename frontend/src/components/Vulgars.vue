<script setup lang="ts">
import { useClipboard } from '@vueuse/core'

const { text, copy, copied } = useClipboard()

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
  <div class="grid gap-2 grid-cols-8">
    <button
      v-for="(vulgarFraction, index) in vulgarFractions" :key="index"
      class="textButton w-auto"
      :class="{ 'ring-2 ring-green-500': copied && text === vulgarFraction }"
      @click="copy(vulgarFraction)"
    >
      {{ vulgarFraction }}
    </button>
  </div>
</template>
