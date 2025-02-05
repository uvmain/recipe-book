<script setup lang="ts">
const props = defineProps({
  modelValue: {
    type: Boolean as PropType<boolean | null>,
    default: false,
  },
  label: { type: String, required: true },
})

const emit = defineEmits(['update:modelValue'])

const model = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emit('update:modelValue', value)
  },
})
</script>

<template>
  <div class="w-full flex flex-row gap-4">
    <label :for="label">{{ label }}</label>
    <input
      :id="label"
      v-model="model"
      type="checkbox"
      class="border-1 border-white border-dashed p-1"
    >
  </div>
</template>
