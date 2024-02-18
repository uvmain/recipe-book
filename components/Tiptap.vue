<script setup lang="ts">
import { BubbleMenu, Editor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'

const props = defineProps({
  modelValue: { type: String, required: true, default: null },
  label: { type: String, required: true },
})

const emits = defineEmits(['update:modelValue'])

const editor = ref()

onMounted(() => {
  editor.value = new Editor({
    content: props.modelValue,
    extensions: [
      StarterKit,
    ],
    onUpdate: () => {
      emits('update:modelValue', editor.value.getHTML())
    },
  })
})

onUnmounted(() => {
  editor.value.destroy()
})
</script>

<template>
  <div>
    <label class="m-2 block text-white">
      {{ label }}
    </label>
    <BubbleMenu
      v-if="editor"
      :editor="editor"
      :tippy-options="{ duration: 100 }"
      class="flex"
    >
      <button :class="{ 'is-active': editor.isActive('bold') }" class="tip-tap-button" @click="editor.chain().focus().toggleBold().run()">
        bold
      </button>
      <button :class="{ 'is-active': editor.isActive('italic') }" class="tip-tap-button" @click="editor.chain().focus().toggleItalic().run()">
        italic
      </button>
      <button :class="{ 'is-active': editor.isActive('strike') }" class="tip-tap-button" @click="editor.chain().focus().toggleStrike().run()">
        strike
      </button>
    </BubbleMenu>
    <EditorContent id="editor" :editor="editor" class="add-form-component" />
  </div>
</template>
