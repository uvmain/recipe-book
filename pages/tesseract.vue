<script setup lang="ts">
import { PSM, createWorker } from 'tesseract.js'

const reognisedText = ref('')

async function recognize() {
  const worker = await createWorker('eng')

  await worker.setParameters({
    tessedit_pageseg_mode: PSM.AUTO_OSD,
  })

  const img = document.getElementById('text-img') as HTMLImageElement

  if (img) {
    const { data: { text } } = await worker.recognize(img)
    reognisedText.value = text
  }

  await worker.terminate()
}
</script>

<template>
  <div class="font-sans antialiased text-center text-bluegray-700 mt-60">
    <button @click="recognize">
      recognize
    </button>
    <img id="text-img" alt="Vue logo" src="/ocr-test.jpg">
    <br>
    <div v-if="reognisedText">
      {{ reognisedText }}
    </div>
  </div>
</template>
