<script setup lang="ts">
defineProps({
  modelValue: { type: null, required: true, default: null },
  label: { type: String, default: undefined },
  placeholder: { type: String, default: 'Please select a value' },
  options: { type: Array as PropType<string[]>, required: true },
})

defineEmits(['update:modelValue'])
const select = ref<HTMLSelectElement>()
</script>

<template>
  <div class="w-full flex flex-row gap-4">
    <label :for="label">{{ label }}</label>
    <select
      ref="select"
      class="mb-6 max-w-xs w-full border-1 p-2 focus:ring-3"
      @change="$emit('update:modelValue', select?.value)"
    >
      <option
        :selected="(modelValue === null) || (modelValue === undefined) || (modelValue === '')"
        disabled hidden
      >
        {{ placeholder }}
      </option>

      <option
        v-for="(option, index) of options" :key="index"
        :selected="option === `${modelValue}`"
        :value="option"
      >
        {{ option }}
      </option>
    </select>
  </div>
</template>
