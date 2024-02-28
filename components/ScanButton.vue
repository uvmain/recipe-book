<script setup lang="ts">
const props = defineProps({
  download: { type: Boolean, default: false },
})

defineEmits(['add', 'reset', 'download'])

const columns = computed(() => {
  if (props.download)
    return 'grid-cols-3'
  return 'grid-cols-2'
})
</script>

<template>
  <div
    class="grid gap-2 items-baseline"
    :class="columns"
  >
    <button class="add-form-component" @click="$emit('add')">
      <Icon name="lucide:plus" class="scale-140" />
    </button>
    <button class="add-form-component" @click="$emit('reset')">
      <Icon name="carbon:reset" class="scale-140" />
    </button>
    <button v-if="download" class="add-form-component" @click="$emit('download')">
      <Icon name="material-symbols:download" class="scale-140" />
    </button>
  </div>
</template>
