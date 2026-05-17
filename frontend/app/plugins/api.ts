export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()
  const baseURL = import.meta.server ? config.apiBase : config.public.apiBase

  const api = $fetch.create({
    baseURL,
  })

  return {
    provide: { api },
  }
})
