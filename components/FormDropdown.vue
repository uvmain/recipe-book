<script setup lang="ts">
export interface DropdownOption {
  title: string
  value: string
}

defineProps({
  modelValue: { type: null, required: true, default: null },
  label: { type: String, required: true },
  options: { type: Array as PropType<DropdownOption[]>, required: true },
})

defineEmits(['update:modelValue'])
const select = ref<HTMLSelectElement>()
</script>

<template>
  <div>
    <label class="m-2 block text-white">
      {{ label }}
    </label>
    <select
      ref="select"
      class="block w-full px-3 py-3 text-base font-normal text-dark bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0"
      @change="$emit('update:modelValue', select?.value)"
    >
      <option
        :selected="(modelValue === null) || (modelValue === undefined)"
      >
        Please select...
      </option>

      <option
        v-for="(option, index) of options" :key="index"
        :selected="option.value === modelValue"
        :value="option.value"
      >
        {{ option.title }}
      </option>
    </select>
  </div>
</template>
