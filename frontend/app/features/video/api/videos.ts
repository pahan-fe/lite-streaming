import { VideoListSchema } from '../schemas/video.schema'

export const fetchVideoList = async () => {
  const { $api } = useNuxtApp()

  const raw = await $api('/api/videos')

  return VideoListSchema.parse(raw)
}
