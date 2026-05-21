import { uploadVideo } from '../api/videos'

export const useUploadVideo = () => {
  const isPending = ref(false)
  const error = ref<Error | null>(null)

  const upload = async (file: File) => {
    isPending.value = true
    error.value = null

    try {
      await uploadVideo(file)
      await refreshNuxtData('videos')
      await navigateTo('/', { replace: true })
    } catch (err) {
      error.value = err instanceof Error ? err : new Error('Unknown error')
    } finally {
      isPending.value = false
    }
  }

  return { upload, isPending: readonly(isPending), error: readonly(error) }
}
