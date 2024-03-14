<script lang="ts" setup>
const props = defineProps({
  modelValue: { type: Boolean, default: false },
  label: { type: String, default: '' },
  id: { type: String, default: '' },
  boxColour: { type: String, default: 'text-white' },
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

function nextValue(value: boolean) {
  return Object.is(value, null) ? true : !value
}

function handleKeyEvent(e: Event) {
  emit('update:modelValue', nextValue(props.modelValue))
  e.preventDefault()
}
</script>

<template>
  <div>
    <label class="text-zinc-800 my-2 block min-w-1/4" :for="id">
      {{ label }}
    </label>
    <input :id="id" v-model="model" type="checkbox" :value="modelValue" class="hidden">
    <div
      class="check-box"
      tabindex="0"
      role="checkbox"
      @click="emit('update:modelValue', nextValue(props.modelValue))"
      @keydown.space="handleKeyEvent"
    >
      <span class="text-2xl" :class="boxColour">
        <Icon v-if="Object.is(modelValue, null)" name="ri:checkbox-indeterminate-fill" />
        <Icon v-else-if="modelValue" name="ri:checkbox-fill" />
        <Icon v-else name="ri:checkbox-blank-line" />
      </span>
    </div>
  </div>
</template>
