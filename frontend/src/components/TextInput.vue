<script setup lang="ts">
const props = defineProps({
  modelValue: { type: String, required: true },
  label: { type: String, required: true },
  type: { type: String, required: false, default: 'text' },
})

const emits = defineEmits(['update:modelValue'])

const model = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emits('update:modelValue', value)
  },
})

function markdownText() {
  model.value = `- ${model.value}`.replaceAll('\n', '\n- ')
}

function handleInput(event: any) {
  model.value = event.target.value
}
</script>

<template>
  <div class="relative">
    <div class="w-full flex flex-col gap-2">
      <label :for="label">{{ label }}</label>
      <input
        v-if="type === 'text'"
        :id="label"
        :placeholder="model"
        :value="model"
        :type="type"
        class="w-full p-2 rounded text border recipeCardBackground"
        @input="handleInput"
      />
      <textarea
        v-if="type === 'textarea'"
        :id="label"
        :placeholder="model"
        :value="model"
        class="w-full p-2 rounded min-h-40 text border recipeCardBackground"
        @input="handleInput"
      />
    </div>
    <div v-if="type === 'textarea'" class="absolute flex flex-col gap-2 -right-20 top-2">
      <button
        id="markdownList"
        class="textButton w-auto"
        @click="markdownText()"
      >
        .md listify
      </button>
    </div>
  </div>
</template>
