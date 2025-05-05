<script setup lang="ts">
defineProps({
  modelValue: { type: null, required: false, default: null },
  label: { type: String, default: undefined },
  placeholder: { type: String, default: 'Please select a value' },
  options: { type: Array as PropType<string[]>, required: true },
})

defineEmits(['update:modelValue'])
const select = ref<HTMLSelectElement>()
</script>

<template>
  <div class="flex flex-col w-full gap-2">
    <label :for="label">{{ label }}</label>
    <select
      ref="select"
      class="w-full border-1 max-w-xs p-2 focus:ring-3 rounded"
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
