import type { SessionCheck } from '../types/auth'
import { useSessionStorage } from '@vueuse/core'
import { backendFetchRequest } from './fetchFromBackend'

const userLoginState = useSessionStorage('untrustedLoginState', false)

export async function checkIfLoggedIn(): Promise<boolean> {
  try {
    const response = await backendFetchRequest('check-session', {
      method: 'GET',
      credentials: 'include',
    })
    const jsonData = await response.json() as SessionCheck
    userLoginState.value = jsonData.loggedIn
    return jsonData.loggedIn
  }
  catch {
    userLoginState.value = false
    return false
  }
}
