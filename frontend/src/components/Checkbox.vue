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
  <div class="flex w-full flex-row gap-4">
    <label :for="label">{{ label }}</label>
    <input
      :id="label"
      v-model="model"
      type="checkbox"
      class="accent-slate-600 dark:accent-slate-400"
    >
  </div>
</template>
