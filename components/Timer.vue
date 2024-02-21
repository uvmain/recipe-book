<script setup lang="ts">
const props = defineProps({
  minutes: { type: Number, default: 1 },
})

const seconds = ref<number>()
const counting = ref(false)

seconds.value = props.minutes * 60

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
}

const timeLeft = computed(() => {
  if (typeof seconds.value !== 'number' || seconds.value < 0) {
    return 'Invalid input'
  }
  const hours = Math.floor(seconds.value / 3600)
  const minutes = Math.floor((seconds.value % 3600) / 60)
  const remainingSeconds = seconds.value % 60
  return `${String(hours).padStart(2, '0')} : ${String(minutes).padStart(2, '0')} : ${String(remainingSeconds).padStart(2, '0')}`
})

const label = computed(() => {
  const hours = Math.floor(props.minutes / 60)
  const minutes = Math.floor(props.minutes % 60)
  const hoursLabel = hours > 1 ? `${hours} hours` : hours === 1 ? `${hours} hour` : ''
  const minutesLabel = minutes > 1 ? `${minutes} minutes` : minutes === 1 ? `${minutes} minute` : ''
  return `${hoursLabel} ${minutesLabel}`
})
</script>

<template>
  <div>
    <label class="text-white m-2 block">
      {{ label }}
    </label>
    <button
      type="button"
      class="text-white text-xl rounded-lg font-medium md:text-3xl px-5 py-2.5 me-2 mb-2 text-center"
      :class="[counting ? 'bg-green-600' : seconds === 0 ? 'bg-red-800' : 'bg-cyan-800']"
      @click="toggleCounting"
    >
      <div>
        {{ timeLeft }}
      </div>
    </button>
    <button
      type="button"
      class="text-white text-xl rounded-lg font-medium md:text-3xl px-5 py-2.5 me-2 mb-2 text-center bg-teal-800"
      @click="resetTimer"
    >
      <Icon name="material-symbols:device-reset" />
    </button>
  </div>
</template>
