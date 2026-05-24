import { deleteVideo } from '../api/videos'

export const useDeleteVideo = () => {
  const isPending = ref(false)
  const error = ref<Error | null>(null)

  const deleteById = async (id: string) => {
    isPending.value = true
    error.value = null

    try {
      await deleteVideo(id)
      await refreshNuxtData('videos')
    } catch (err) {
      error.value = err instanceof Error ? err : new Error('Unknown error')
      throw err
    } finally {
      isPending.value = false
    }
  }

  return { deleteById, isPending: readonly(isPending), error: readonly(error) }
}
