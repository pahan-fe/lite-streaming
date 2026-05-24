import { fetchVideoById } from '../api/videos'

export const useVideo = (id: string) => useAsyncData(`video-${id}`, () => fetchVideoById(id))
