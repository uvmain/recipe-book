<script setup lang="ts">
defineProps({
  id: { type: String, required: true },
  label: { type: String, required: true, default: '' },
  modelValue: { type: String, default: '' },
  required: { type: Boolean, default: false },
})
const emit = defineEmits(['update:modelValue'])

function handleInput(e: Event) {
  const value = (e.target as HTMLInputElement).value
  emit('update:modelValue', value)
}
</script>

<template>
  <div class="pl-5">
    <label class="m-2 block text-white" :for="id">
      {{ label }}
    </label>
    <textarea
      :id="id"
      :value="modelValue"
      class="block w-full px-3 py-3 text-base font-normal text-dark bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0"
      :required="required"
      @input="handleInput"
    />
  </div>
</template>
