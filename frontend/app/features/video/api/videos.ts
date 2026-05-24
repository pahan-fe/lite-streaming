import { VideoListSchema, VideoSchema } from '../schemas/video.schema'

export const fetchVideoList = async () => {
  const { $api } = useNuxtApp()

  const raw = await $api('/api/videos')

  return VideoListSchema.parse(raw)
}

export const uploadVideo = (file: File) => {
  const { $api } = useNuxtApp()

  const formData = new FormData()
  formData.append('file', file)

  return $api('/api/videos', {
    method: 'POST',
    body: formData,
  })
}

export const fetchVideoById = async (id: string) => {
  const { $api } = useNuxtApp()

  const response = await $api(`/api/videos/${id}`)

  return VideoSchema.parse(response)
}
