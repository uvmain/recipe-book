export function debounce<TArgs extends unknown[]>(callback: (...args: TArgs) => void, delay: number) {
  let timer: ReturnType<typeof setTimeout>
  return function (...args: TArgs) {
    clearTimeout(timer)
    timer = setTimeout(() => {
      callback(...args)
    }, delay)
  }
}
