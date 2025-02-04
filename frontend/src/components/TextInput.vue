<script setup lang="ts">
defineProps({
  modelValue: { type: String, required: true },
  label: { type: String, required: true },
  type: { type: String, required: false, default: 'text' },
})

const emits = defineEmits(['update:modelValue'])

function handleInput(event: any) {
  emits('update:modelValue', event.target.value)
}
</script>

<template>
  <div class="w-full flex flex-row gap-4">
    <label :for="label">{{ label }}</label>
    <input
      :id="label"
      :placeholder="modelValue"
      :value="modelValue"
      :type="type"
      class="w-full border-1 border-white border-dashed p-1"
      @input="handleInput"
    >
  </div>
</template>
