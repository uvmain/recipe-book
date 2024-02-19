<script setup lang="ts">
import { Editor, EditorContent } from '@tiptap/vue-3'
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
    <div v-if="editor" class="flex col-auto gap-1 my-1">
      <button :class="{ 'is-active': editor.isActive('bold') }" class="tip-tap-button" @click="editor.chain().focus().toggleBold().run()">
        <Icon name="material-symbols:format-bold" />
      </button>
      <button :class="{ 'is-active': editor.isActive('italic') }" class="tip-tap-button" @click="editor.chain().focus().toggleItalic().run()">
        <Icon name="material-symbols:format-italic" />
      </button>
      <button :class="{ 'is-active': editor.isActive('hardBreak') }" class="tip-tap-button" @click="editor.chain().focus().setHardBreak().run()">
        <Icon name="material-symbols:keyboard-return" />
      </button>
      <button :class="{ 'is-active': editor.isActive('bulletList') }" class="tip-tap-button" @click="editor.chain().focus().toggleBulletList().run()">
        <Icon name="material-symbols:format-list-bulleted" />
      </button>
      <button :class="{ 'is-active': editor.isActive('listItem') }" class="tip-tap-button" @click="editor.chain().focus().toggleOrderedList().run()">
        <Icon name="material-symbols:format-list-numbered" />
      </button>
      <button :class="{ 'is-active': editor.isActive('strike') }" class="tip-tap-button" @click="editor.chain().focus().toggleStrike().run()">
        <Icon name="material-symbols:format-strikethrough" />
      </button>
      <button :class="{ 'is-active': editor.isActive('strike') }" class="tip-tap-button" @click="editor.chain().focus().setHorizontalRule().run()">
        <Icon name="material-symbols:horizontal-rule" />
      </button>
    </div>
    <EditorContent id="editor" :editor="editor" class="add-form-component" />
  </div>
</template>
