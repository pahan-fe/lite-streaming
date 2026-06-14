import { VideoListResponseSchema, VideoSchema } from '../schemas/video.schema'

interface FetchVideoListQuery {
  limit?: number
  cursor?: string
}
export const fetchVideoList = async ({ limit, cursor }: FetchVideoListQuery = {}) => {
  const { $api } = useNuxtApp()

  const query: FetchVideoListQuery = {}
  if (cursor) {
    query.cursor = cursor
  }
  if (limit) {
    query.limit = limit
  }

  const raw = await $api('/api/videos', { query })

  return VideoListResponseSchema.parse(raw)
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

export const deleteVideo = (id: string) => {
  const { $api } = useNuxtApp()

  return $api(`/api/videos/${id}`, {
    method: 'DELETE',
  })
}
