<script setup lang="ts">
const props = defineProps({
  minutes: { type: Number, default: 1 },
})

const seconds = ref<number>()
const counting = ref(false)

seconds.value = props.minutes * 60

const alarm = new Audio('/audio/sonnette_reveil.wav')

const interval = setInterval(() => {
  if (seconds.value === 0) {
    clearInterval(interval)
    counting.value = false
  }
  else if (seconds.value && counting.value) {
    seconds.value -= 1
  }
  else {
    counting.value = false
  }
}, 1000)

function toggleCounting() {
  counting.value = !counting.value
}

function resetTimer() {
  counting.value = false
  seconds.value = props.minutes * 60
  alarm.pause()
  alarm.currentTime = 0
}

watch(seconds, (newValue, oldValue) => {
  if (newValue === 0 && oldValue !== 0) {
    alarm.play()
  }
})

const timeLeft = computed(() => {
  if (typeof seconds.value !== 'number' || seconds.value < 0) {
    return 'Invalid input'
  }
  const hours = Math.floor(seconds.value / 3600)
  const minutes = Math.floor((seconds.value % 3600) / 60)
  const remainingSeconds = seconds.value % 60
  return `${String(hours).padStart(2, '0')} : ${String(minutes).padStart(2, '0')} : ${String(remainingSeconds).padStart(2, '0')}`
})

const timerColour = computed(() => {
  if (counting.value)
    return 'border-emerald-600 border-2'
  else {
    if (seconds.value === 0)
      return 'border-red-800 border-2'
    if (seconds.value === props.minutes * 60)
      return 'border-neutral-400'
    else
      return 'border-orange-400 border-2'
  }
})

onUnmounted(() => {
  clearInterval(interval)
})
</script>

<template>
  <div class="flex flex-row gap-1">
    <button
      type="button"
      class="text-lg rounded-lg text-center mb-2 px-5 border-1 border-solid bg-blue-gray-200 text-dark whitespace-nowrap"
      :class="timerColour"
      @click="toggleCounting"
    >
      <div>
        {{ timeLeft }}
      </div>
    </button>
    <button
      type="button"
      class="text-lg rounded-lg px-5 py-2.5 mb-2 text-center border-1 border-solid border-gray-400 bg-blue-gray-300 text-dark"
      @click="resetTimer"
    >
      <Icon name="lucide:rotate-ccw" class="pb-1"/></button>
  </div>
</template>
