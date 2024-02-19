<script setup lang="ts">
defineProps({
  id: { type: String, required: true },
  label: { type: String, required: true, default: '' },
  modelValue: { type: String, default: '' },
  required: { type: Boolean, default: false },
  type: { type: String, default: 'text' },
})
const emit = defineEmits(['update:modelValue'])

function handleInput(e: Event) {
  const value = (e.target as HTMLInputElement).value
  emit('update:modelValue', value)
}
</script>

<template>
  <div>
    <label class="m-2 block text-white" :for="id">
      {{ label }}
    </label>
    <input
      :id="id"
      :value="modelValue"
      class="add-form-component"
      :type="type"
      :required="required"
      @input="handleInput"
    >
  </div>
</template>
