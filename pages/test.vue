<script setup lang="ts">
const { body } = await $fetch('/api/converter', {
  method: 'post',
  body: { test: 123 },
})
</script>

<template>
  <div class="text-black">
    {{ body }}
  </div>
</template>
