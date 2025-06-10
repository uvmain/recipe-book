<script setup lang="ts">
const props = defineProps({
  resizedImage: { type: String, required: true },
  resizedImageWidth: { type: Number, required: true },
  resizedImageHeight: { type: Number, required: true },
  label: { type: String, required: true },
  recipe: { type: Object, required: true },
})

const emit = defineEmits(['update:modelValue', 'update:resizedImage', 'update:resizedImageWidth', 'update:resizedImageHeight'])
const resizedImageUrl = ref<string>('')

const maxPixelSize = ref<number>(800)
const imageQuality = ref(0.9)

const model = computed({
  get() {
    return props.resizedImage
  },
  set(value) {
    emit('update:resizedImage', value)
  },
})

const imageUrlInput = ref<string>('')

function onFileChange(event: Event) {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = () => {
      model.value = reader.result as string
      imageUrlInput.value = ''
    }
    reader.readAsDataURL(file)
  }
}

async function onUrlChange() {
  try {
    const response = await fetch(imageUrlInput.value)
    if (!response.ok)
      throw new Error('Network response was not ok')
    const blob = await response.blob()
    const reader = new FileReader()
    reader.onloadend = () => {
      model.value = reader.result as string
    }
    reader.readAsDataURL(blob)
  }
  catch (error) {
    console.error(`Failed to load image: ${error}`)
  }
}

function onPaste(event: ClipboardEvent) {
  const items = event.clipboardData?.items
  if (items) {
    for (const item of Array.from(items)) {
      if (item.type.startsWith('image/')) {
        const file = item.getAsFile()
        if (file) {
          const reader = new FileReader()
          reader.onload = () => {
            model.value = reader.result as string
            imageUrlInput.value = ''
          }
          reader.readAsDataURL(file)
        }
      }
    }
  }
}

function resizeAndConvertImage(base64Image: string) {
  const img = new Image()
  img.src = base64Image

  img.onload = () => {
    const canvas = document.createElement('canvas')
    const ctx = canvas.getContext('2d')

    if (!ctx) {
      console.error('Could not get canvas context')
      return
    }

    const maxWidth = maxPixelSize.value
    const maxHeight = maxPixelSize.value
    let width = img.width
    let height = img.height

    if (width > height) {
      if (width > maxWidth) {
        height *= maxWidth / width
        width = maxWidth
      }
    }
    else {
      if (height > maxHeight) {
        width *= maxHeight / height
        height = maxHeight
      }
    }

    width = Math.round(width)
    height = Math.round(height)

    canvas.width = width
    canvas.height = height
    ctx.drawImage(img, 0, 0, width, height)
    resizedImageUrl.value = canvas.toDataURL('image/webp', imageQuality.value)
    emit('update:resizedImage', resizedImageUrl.value)
    emit('update:resizedImageWidth', width)
    emit('update:resizedImageHeight', height)
  }

  img.onerror = (error) => {
    console.error('Failed to load image', error)
  }
}

const fallbackImageAddress = computed(() => {
  return props.recipe.imageFilename ? `/api/images/${props.recipe.imageFilename}` : '/default.webp'
})

watch(model, () => {
  resizeAndConvertImage(model.value)
})
</script>

<template>
  <div class="flex flex-col items-center text p-6 gap-10 max-w-100vw" @paste="onPaste">
    <label :for="label">{{ label }}</label>
    <div :id="label" class="flex flex-col gap-0">
      <input
        id="fileInput"
        type="file"
        accept="image/*"
        class="hidden"
        @change="onFileChange"
      />
      <label
        for="fileInput"
        class="border-1 border-solid p-2 text-center block text-sm text-gray-700 border-gray-300 rounded-lg cursor-pointer bg-gray-50 focus:outline-none hover:bg-gray-200 w-28rem"
      >
        Browse
      </label>
      <span class="text-sm">
        or
      </span>
      <input
        v-model="imageUrlInput"
        placeholder="Enter image URL"
        class="block w-28rem p-2 border-1 border-solid border-gray-300 rounded-lg focus:outline-none mt-2 focus:ring-2 focus:ring-blue-600"
        @change="onUrlChange"
      />
      <span class="text-sm">
        or
      </span>
      <textarea
        placeholder="Paste an image here"
        class="block w-28rem mt-2 p-2 border-1 border-solid border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
      ></textarea>
      <div v-if="model" class="mt-4">
        <img
          :src="model"
          alt="Uploaded Image"
          class="rounded-lg w-28rem shadow-md object-contain"
        />
      </div>
      <div v-else-if="recipe.imageFilename" class="mt-4">
        <img
          :src="fallbackImageAddress"
          alt="Uploaded Image"
          class="rounded-lg shadow-md w-28rem object-contain"
        />
      </div>
    </div>
    <div class="flex flex-col gap-4 mt-6">
      <label>
        Max Pixel Size: {{ maxPixelSize }}
        <input
          v-model="maxPixelSize"
          type="range"
          min="100"
          max="2000"
          step="100"
          class="w-full"
        />
      </label>
      <label>
        Image Quality: {{ imageQuality }}
        <input
          v-model="imageQuality"
          type="range"
          min="0.1"
          max="1"
          step="0.1"
          class="w-full"
        />
      </label>
      <button
        class="p-4 h-auto w-auto headerButton"
        @click="resizeAndConvertImage(model)"
      >
        Resize Image
      </button>
    </div>
  </div>
</template>
