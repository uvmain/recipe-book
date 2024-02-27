<script setup lang="ts">
defineEmits(['add', 'reset'])
</script>

<template>
  <div class="grid grid-cols-2 grid-rows-none gap-2 items-baseline">
    <button class="add-form-component" @click="$emit('add')">
      <Icon name="lucide:plus" class="scale-140" />
    </button>
    <button class="add-form-component" @click="$emit('reset')">
      <Icon name="carbon:reset" class="scale-140" />
    </button>
  </div>
</template>
