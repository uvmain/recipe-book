<script setup lang="ts">
import type { SessionCheck } from '../types/auth'
import { onClickOutside, useSessionStorage } from '@vueuse/core'
import { checkIfLoggedIn } from '../composables/auth'
import { backendFetchRequest } from '../composables/fetchFromBackend'

defineProps({
  isOpen: Boolean,
})

const emits = defineEmits(['modalClose'])

const username = ref('')
const password = ref('')
const target = ref(null)
const userLoginState = useSessionStorage('untrustedLoginState', false)

async function login() {
  const formData = new FormData()
  formData.append('username', username.value)
  formData.append('password', password.value)

  const response = await backendFetchRequest('login', {
    body: formData,
    method: 'POST',
  })
  const jsonData = await response.json() as SessionCheck
  userLoginState.value = jsonData.loggedIn
  emits('modalClose')
}

function cancel() {
  emits('modalClose')
}

async function logout() {
  await backendFetchRequest('logout', {
    method: 'GET',
    credentials: 'include',
  })
  userLoginState.value = false
  emits('modalClose')
}

onBeforeMount(() => {
  checkIfLoggedIn()
})

onClickOutside(target, () => emits('modalClose'))
</script>

<template>
  <div v-if="isOpen" class="top-0 text fixed left-0 z-999 size-full backdrop-blur-xl">
    <div v-if="!userLoginState" @keydown.escape="cancel">
      <div ref="target" class="mx-auto mt-150px p-30px modal recipeCardBackground border rounded-md max-w-80vw lg:max-w-[min(400px,40vw)]">
        <div class="flex flex-col gap-4 p-6">
          <form class="flex flex-col gap-4 items-center">
            <div class="flex flex-row items-center gap-2">
              <label for="username">Username:</label>
              <input id="username" v-model="username" type="text" name="username" autocomplete="username" class="text border recipeCardBackground">
            </div>
            <div class="flex flex-row items-center gap-2">
              <label for="password">Password:</label>
              <input id="password" v-model="password" type="password" name="password" autocomplete="current-password" class="text border recipeCardBackground" @keydown.enter="login">
            </div>
          </form>
        </div>
        <div class="flex justify-center gap-4">
          <button aria-label="cancel" class="textButton w-auto" @click="cancel">
            Cancel
          </button>
          <button aria-label="login" class="textButton w-auto" @click="login">
            Login
          </button>
        </div>
      </div>
    </div>
    <div v-else @keydown.escape="cancel">
      <div class="mx-auto mb-auto mt-150px w-300px px-30px pb-30px pt-20px modal recipeCardBackground border rounded-md">
        <div class="text-center mb-2 py-4">
          You are logged in.
        </div>
        <div class="flex justify-center gap-4">
          <button aria-label="cancel" class="textButton w-auto" @click="cancel">
            Cancel
          </button>
          <button aria-label="logout" class="textButton w-auto" @click="logout">
            Logout
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
