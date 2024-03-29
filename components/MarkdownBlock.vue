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
    arrayItem = arrayItem.replaceAll('1/5', '⅕').replaceAll('1/4', '¼').replaceAll('1/3', '⅓').replaceAll('1/2', '½').replaceAll('2/3', '⅔').replaceAll('3/4', '¾')
    arrayItem = arrayItem.replaceAll('puree', 'purée')
    arrayItem = arrayItem.replaceAll('saute', 'sauté')

    let tag = ''
    let step: string | null = null
    if (arrayItem.startsWith('- ')) {
      tag = 'li'
      step = arrayItem.replace('- ', '')
    }
    else if (arrayItem.startsWith('**')) {
      tag = 'str'
      step = arrayItem.replaceAll('*', '')
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

    const timers = step ? getTimer(step) : []
    for (const timer of timers) {
      parsedMd.push({ tag: 'tm', step: `${timer}` })
    }
  })
  return parsedMd
}

const mdItems = computed(() => {
  return parsedMarkdown(props.markdownString)
})

function getStepWithLink(step: string) {
  // return step.split(/(<)/)
  const splitStep = step.split('<a href="')
  const route = splitStep[1].split('"')[0]
  const linkText = splitStep[1].split('>')[1].split('<')[0]
  const remainingText = splitStep[1].split('</a>')[1]
  return {
    startingText: splitStep[0],
    route,
    linkText,
    remainingText,
  }
}
</script>

<template>
  <div class="rounded-lg pt-1 p-2 pr-4" :class="propsClass">
    <h3 class="font-bold text-xl pl-2">
      {{ label }}
    </h3>
    <div v-for="(mdItem, index) of mdItems" :key="index" class="leading-5 ml-4">
      <ul v-if="mdItem.tag === 'li'" class="my-2">
        <li v-if="mdItem.step?.includes('<a')">
          {{ getStepWithLink(mdItem.step).startingText }}<a
            :href="getStepWithLink(mdItem.step).route"
            class="text-white underline-none"
          >{{ getStepWithLink(mdItem.step).linkText }}</a>{{
            getStepWithLink(mdItem.step).remainingText }}
        </li>
        <li v-else v-text="mdItem.step" />
      </ul>
      <strong v-if="mdItem.tag === 'str'" v-text="mdItem.step" />
      <br v-if="mdItem.tag === 'br'">
      <p v-if="mdItem.tag === 'p'" />
      <div v-if="mdItem.tag === 'sp'">
        <span v-text="mdItem.step" />
      </div>
      <Timer v-if="mdItem.step && mdItem.tag === 'tm'" :minutes="Number(mdItem.step)" class="ml-8 mt-4" />
    </div>
  </div>
</template>
