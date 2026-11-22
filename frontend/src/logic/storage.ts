import type { Ref, UnwrapRef } from 'vue'

export function useSessionStorage<T>(key: string, defaultValue: T): Ref<UnwrapRef<T>, T | UnwrapRef<T>> {
  const storedValue = sessionStorage.getItem(key)
  const data = ref<T>((storedValue != null)
    ? JSON.parse(storedValue) as T
    : defaultValue)

  watch(data, (newValue) => {
    sessionStorage.setItem(key, JSON.stringify(newValue))
  }, { deep: true })

  return data as Ref<UnwrapRef<T>, T | UnwrapRef<T>>
}

export function useLocalStorage<T>(key: string, defaultValue: T): Ref<UnwrapRef<T>, T | UnwrapRef<T>> {
  const storedValue = localStorage.getItem(key)
  const data = ref<T>((storedValue != null)
    ? JSON.parse(storedValue) as T
    : defaultValue)

  watch(data, (newValue) => {
    localStorage.setItem(key, JSON.stringify(newValue))
  }, { deep: true })

  return data as Ref<UnwrapRef<T>, T | UnwrapRef<T>>
}
