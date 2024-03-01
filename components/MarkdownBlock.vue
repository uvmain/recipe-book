<script setup lang="ts">
const props = defineProps({
  label: { type: String, required: true },
  markdownString: { type: String, required: true },
  propsClass: { type: String, default: 'bg-gray-600' },
})

interface parsedMdItem {
  tag: string
  step: string | null
}

function parsedMarkdown(markdownItem: string): parsedMdItem[] {
  const parsedMd: parsedMdItem[] = []
  const initialArray: string[] = markdownItem.replaceAll('\n\n', '\n<br>\n').split('\n')
  initialArray.forEach((arrayItem: string) => {
    let tag = ''
    let step: string | null = null
    if (arrayItem.startsWith('- ')) {
      tag = 'li'
      step = arrayItem.replace('- ', '')
    }
    else if (arrayItem.startsWith('* * *') || arrayItem.startsWith('***')) {
      tag = 'br'
      step = null
    }
    else if (arrayItem.startsWith('**')) {
      tag = 'str'
      step = arrayItem.replace('* ', '')
    }
    else if (arrayItem === '<br>') {
      tag = 'br'
      step = null
    }
    else if (arrayItem.length === 0) {
      tag = 'p'
      step = null
    }
    else if (arrayItem !== '\\') {
      tag = 'sp'
      step = arrayItem
    }
    parsedMd.push({ tag, step })

    const timer = step ? getTimer(step) : null
    if (timer) {
      console.log(timer)
      parsedMd.push({ tag: 'tm', step: `${timer}` })
    }
  })
  return parsedMd
}

const mdItems = computed(() => {
  return parsedMarkdown(props.markdownString)
})
</script>

<template>
  <div class="rounded-lg p-2 pt-1" :class="propsClass">
    <h3 class="text-xl font-bold ml-2">
      {{ label }}
    </h3>
    <div v-for="(mdItem, index) of mdItems" :key="index" class="pl-2 md:pl-4 leading-relaxed">
      <ul v-if="mdItem.tag === 'li'" class="my-2">
        <li v-text="mdItem.step" />
      </ul>
      <strong v-if="mdItem.tag === 'str'" v-text="mdItem.step" />
      <br v-if="mdItem.tag === 'br'">
      <p v-if="mdItem.tag === 'p'" />
      <span v-if="mdItem.tag === 'sp'" v-text="mdItem.step" />
      <Timer v-if="mdItem.step && mdItem.tag === 'tm'" :minutes="Number(mdItem.step)" class="ml-8 mt-4" />
    </div>
  </div>
</template>
