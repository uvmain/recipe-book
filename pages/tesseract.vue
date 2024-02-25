<script setup lang="ts">
import { PSM, createWorker } from 'tesseract.js'
import '@recogito/annotorious/dist/annotorious.min.css'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'

const recognizedParagraphs = ref<string[]>([])
const anno = ref<typeof import('@recogito/annotorious')>(null)

const imageToRecognise = ref()
const source = ref('')
const author = ref('')
const title = ref<string>('')
const ingredients = ref('')
const instructions = ref('')

async function recognize() {
  recognizedParagraphs.value = []
  const worker = await createWorker('eng')

  await worker.setParameters({
    tessedit_pageseg_mode: PSM.AUTO_OSD,
    preserve_interword_spaces: '1',
  })
  if (imageToRecognise.value) {
    const { data: { paragraphs } } = await worker.recognize(imageToRecognise.value)
    paragraphs.forEach((paragraph) => {
      for (const newLine of paragraph.text.split('\n'))
        recognizedParagraphs.value.push(newLine)
    })
  }
  await worker.terminate()
}

async function addToTitle() {
  await recognize()
  title.value = title.value + (recognizedParagraphs.value).join(' ')
}

async function addToSource() {
  await recognize()
  source.value = source.value + (recognizedParagraphs.value).join(' ')
}

async function addToAuthor() {
  await recognize()
  author.value = author.value + (recognizedParagraphs.value).join(' ')
}

async function addToIngredients() {
  await recognize()

  let markdown = ''
  const hasNumberRegex = /\d/
  recognizedParagraphs.value.forEach((text: string) => {
    if (markdown.length) {
      markdown += '\n'
    }
    if (!hasNumberRegex.test(text) && text.length > 2) {
      markdown += `**${text}**`
    }
    else if (text.length > 2) {
      markdown += `- ${text}`
    }
  })
  ingredients.value += markdown.replaceAll('\n\n\n', '\n\n')
}

async function addToInstructions() {
  await recognize()
  const startsWithUppercaseOrNumber = /^[A-Z0-9]/
  let markdown = ''
  recognizedParagraphs.value.forEach((text: string) => {
    if (markdown.length && text.length) {
      markdown += '\n'
    }
    if (text.length && startsWithUppercaseOrNumber.test(text)) {
      markdown += `- ${text}`
    }
    else if (text.length) {
      markdown += text
    }
  })
  instructions.value += markdown.replaceAll('\n\n\n', '\n\n')
}

onMounted(async () => {
  const Anno = await import('@recogito/annotorious')
  anno.value = new Anno.Annotorious({
    image: document.getElementById('text-img'),
    allowEmpty: true,
  })
  if (anno.value)
    anno.value.on('createAnnotation', async (annotation: any) => {
      const id = annotation.id
      const snippetObject: { snippet: HTMLCanvasElement } = await anno.value.getImageSnippetById(id)
      const canvas = document.createElement('canvas')
      const ctx = canvas.getContext('2d')
      canvas.width = snippetObject.snippet.width * 2
      canvas.height = snippetObject.snippet.height * 2
      ctx?.drawImage(snippetObject.snippet, 0, 0, canvas.width, canvas.height)
      imageToRecognise.value = canvas.toDataURL()
    })
})
</script>

<template>
  <ClientOnly>
    <div class="font-sans antialiased text-center text-bluegray-700 mt-10">
      <div class="flex flex-row">
        <img id="text-img" alt="Vue logo" src="/ocr-test.jpg" class="mx-4 max-w-full ml-auto border-solid border-gray-300 max-h-3/4" @mousedown.prevent="null">
        <div class="max-w-1/2">
          <div class="ml-4 text-left mb-4 flex gap-2 items-center">
            <button class="p-2 min-w-25" @click="addToSource">
              Add to Source
            </button>
            <button class="p-2 min-w-25" @click="source = ''">
              Reset Source
            </button>
            <span class="font-bold text-white text-3xl ml-4">
              Source:
            </span>
            <input
              v-model="source"
              class="add-form-component"
            >
          </div>
          <div class="ml-4 text-left mb-4 flex gap-2 items-center">
            <button class="p-2 min-w-25" @click="addToAuthor">
              Add to Author
            </button>
            <button class="p-2 min-w-25" @click="author = ''">
              Reset Author
            </button>
            <span class="font-bold text-white text-3xl ml-4">
              Author:
            </span>
            <input
              v-model="author"
              class="add-form-component"
            >
          </div>
          <div class="ml-4 text-left mb-4 flex gap-2 items-center">
            <button class="p-2 min-w-25" @click="addToTitle">
              Add to Title
            </button>
            <button class="p-2 min-w-25" @click="title = ''">
              Reset Title
            </button>
            <span class="font-bold text-white text-3xl ml-4">
              Title:
            </span>
            <input
              v-model="title"
              class="add-form-component"
            >
          </div>
          <div class="ml-4 text-left mb-4">
            <div class="mb-2 flex gap-2">
              <button class="p-2" @click="addToIngredients">
                Add to ingredients
              </button>
              <button class="p-2" @click="ingredients = ''">
                Reset Ingredients
              </button>
              <span v-if="ingredients" class="font-bold text-white text-3xl">
                Ingredients:
              </span>
            </div>
            <MdEditor v-if="ingredients" v-model="ingredients" editor-id="ingredients" class="add-form-component" language="en-US" />
          </div>
          <div class="ml-4 text-left">
            <div class="mb-2 flex gap-2">
              <button class="p-2" @click="addToInstructions">
                Add to Instructions
              </button>
              <button class="ml-1 p-2" @click="instructions = ''">
                Reset Instructions
              </button>
              <span v-if="instructions" class="font-bold text-white text-3xl">
                Instructions:
              </span>
            </div>
            <MdEditor v-if="instructions" v-model="instructions" editor-id="instructions" class="add-form-component" language="en-US" />
          </div>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>
