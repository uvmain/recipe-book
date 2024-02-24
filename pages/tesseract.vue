<script setup lang="ts">
import { PSM, createWorker } from 'tesseract.js'
import '@recogito/annotorious/dist/annotorious.min.css'

const recognizedParagraphs = ref<string[]>([])
const anno = ref<typeof import('@recogito/annotorious')>(null)

const temp = ref()

async function recognize() {
  recognizedParagraphs.value = []
  const worker = await createWorker('eng')

  await worker.setParameters({
    tessedit_pageseg_mode: PSM.AUTO_OSD,
    preserve_interword_spaces: '1',
  })

  const img = document.getElementById('text-img') as HTMLImageElement

  if (img) {
    const { data: { paragraphs } } = await worker.recognize(img)
    paragraphs.forEach((paragraph) => {
      if (paragraph.text.replaceAll('\n', '').length > 2)
        recognizedParagraphs.value.push(paragraph.text)
    })
  }

  await worker.terminate()
}

onMounted(async () => {
  const Anno = await import('@recogito/annotorious')
  anno.value = new Anno.Annotorious({
    image: document.getElementById('text-img'),
  })
  if (anno.value)
    anno.value.on('createAnnotation', async (annotation: any) => {
      const id = annotation.id
      const snippetObject: { snippet: HTMLCanvasElement } = await anno.value.getImageSnippetById(id)
      temp.value = snippetObject.snippet.toDataURL()
    })
})
</script>

<template>
  <ClientOnly>
    <div class="font-sans antialiased text-center text-bluegray-700 mt-10">
      <div class="grid grid-cols-2">
        <img id="text-img" alt="Vue logo" src="/ocr-test.jpg" class="mx-4 max-w-full ml-auto border-solid border-gray-300 max-h-3/4" @mousedown.prevent="null">
        <div class="ml-4 text-left">
          <img :src="temp">
          <button class="p-2 mb-4" @click="recognize">
            recognize
          </button>
          <div v-if="recognizedParagraphs.length" class="text-white">
            <div v-for="(text, index) in recognizedParagraphs" :key="index">
              {{ text }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>
