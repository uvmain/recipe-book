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
    return 'bg-emerald-600'
  else {
    if (seconds.value === 0)
      return 'bg-red-800'
    if (seconds.value === props.minutes * 60)
      return 'bg-neutral-400'
    else
      return 'bg-orange-400'
  }
})

onUnmounted(() => {
  clearInterval(interval)
})
</script>

<template>
  <div class="flex flex-row">
    <button
      type="button"
      class="text-white text-xl rounded-lg text-center mb-2 font-medium md:text-3xl px-5 py-2.5 me-2"
      :class="timerColour"
      @click="toggleCounting"
    >
      <div>
        {{ timeLeft }}
      </div>
    </button>
    <button
      type="button"
      class="text-white text-xl rounded-lg font-medium md:text-3xl px-5 py-2.5 me-2 mb-2 text-center bg-zinc-500"
      @click="resetTimer"
    >
      <Icon name="material-symbols:device-reset" />
    </button>
  </div>
</template>
