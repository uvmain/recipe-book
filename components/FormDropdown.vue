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
    <label class="block text-zinc-800 my-2 min-w-1/4">
      {{ label }}
    </label>
    <select
      ref="select"
      class="add-form-component"
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
