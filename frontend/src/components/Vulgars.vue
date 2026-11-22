<script setup lang="ts">
function useClipboard() {
  const text = ref('')
  const copied = ref(false)

  function copy(value: string) {
    text.value = value
    navigator.clipboard.writeText(value).catch((err) => {
      console.error('Failed to copy text: ', err)
    })
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  }

  return { text, copy, copied }
}

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
