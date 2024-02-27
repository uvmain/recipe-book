<script setup lang="ts">
import { PSM, createWorker } from 'tesseract.js'
import '@recogito/annotorious/dist/annotorious.min.css'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'

const recognizedParagraphs = ref<string[]>([])
const anno = ref<typeof import('@recogito/annotorious')>(null)
const sourceImage = ref()
const imageToRecognise = ref()
const recipeImage = ref()
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

async function handleImageChange(event: Event) {
  anno.value?.destroy()

  const input = event.target as HTMLInputElement
  if (input.files && input.files[0]) {
    const reader = new FileReader()
    reader.onload = (e) => {
      if (e.target && typeof e.target.result === 'string') {
        sourceImage.value = e.target.result
      }
    }
    reader.readAsDataURL(input.files[0])
  }
  const Anno = await import('@recogito/annotorious')
  anno.value = new Anno.Annotorious({
    image: document.getElementById('text-img'),
    disableEditor: true,
  })

  anno.value.on('createSelection', async (annotation: any) => {
    const id = annotation.id
    const snippetObject: { snippet: HTMLCanvasElement } = await anno.value.getImageSnippetById(id)
    updateImageToReconise(snippetObject)
  })
  // eslint-disable-next-line unused-imports/no-unused-vars
  anno.value.on('changeSelected', async (selected: any, previous: any) => {
    const id = selected.annotation.id
    const snippetObject: { snippet: HTMLCanvasElement } = await anno.value.getImageSnippetById(id)
    updateImageToReconise(snippetObject)
  })
  anno.value.on('changeSelectionTarget', async (annotation: any) => {
    const id = annotation.id
    const snippetObject: { snippet: HTMLCanvasElement } = await anno.value.getImageSnippetById(id)
    updateImageToReconise(snippetObject)
  })
}

function updateImageToReconise(snippetObject: { snippet: HTMLCanvasElement }) {
  const canvas = document.createElement('canvas')
  const ctx = canvas.getContext('2d')
  canvas.width = snippetObject.snippet.width * 2
  canvas.height = snippetObject.snippet.height * 2
  ctx?.drawImage(snippetObject.snippet, 0, 0, canvas.width, canvas.height)
  imageToRecognise.value = canvas.toDataURL()
}

async function saveRecipeImage() {
  if (imageToRecognise.value) {
    recipeImage.value = imageToRecognise.value
  }
}

async function saveAsWebP() {
  if (imageToRecognise.value) {
    const image = new Image()
    image.src = imageToRecognise.value

    image.onload = async function () {
      const canvas = document.createElement('canvas')
      const ctx = canvas.getContext('2d')

      // Calculate new width and height while preserving aspect ratio
      let newWidth = image.width
      let newHeight = image.height
      if (image.width > 800 || image.height > 800) {
        if (image.width > image.height) {
          newWidth = 800
          newHeight = Math.round((800 / image.width) * image.height)
        }
        else {
          newHeight = 800
          newWidth = Math.round((800 / image.height) * image.width)
        }
      }

      // Resize the canvas
      canvas.width = newWidth
      canvas.height = newHeight
      ctx?.drawImage(image, 0, 0, newWidth, newHeight)

      // Convert canvas to WebP format with 80% quality
      const webpData = canvas.toDataURL('image/webp', 0.8)

      // Convert base64 to blob
      const blob = await fetch(webpData).then(res => res.blob())

      // Create a download link
      const link = document.createElement('a')
      link.href = URL.createObjectURL(blob)
      link.download = 'image.webp'
      link.click()
    }
  }
}

onBeforeUnmount(() => {
  if (anno.value) {
    anno.value.destroy()
  }
})
</script>

<template>
  <ClientOnly>
    <div class="font-sans antialiased text-center text-bluegray-700 mt-10">
      <div class="grid grid-cols-2 gap-4">
        <div class="grid grid-cols-1 gap-4 ml-4 min-w-4/5">
          <input
            id="addFiles"
            type="file"
            accept="image/*"
            multiple="false"
            @change="handleImageChange"
          >
          <div>
            <img v-if="sourceImage" id="text-img" alt="Vue logo" :src="sourceImage" class="w-full h-full" @mousedown.prevent="null">
          </div>
        </div>
        <div class="mr-auto">
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
          <div>
            <button class="p-2 min-w-25" @click="saveRecipeImage">
              Save recipe image
            </button>
            <button class="p-2 min-w-25" @click="saveAsWebP">
              Download image as .webp
            </button>
            <div class="mt-4 max-w-1/4">
              <img v-if="recipeImage" :src="recipeImage" class="w-full h-full" @mousedown.prevent="null">
            </div>
          </div>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>
