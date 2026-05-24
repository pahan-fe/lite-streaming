import { fetchVideoList } from '../api/videos'

import type { Video } from '../schemas/video.schema'

export const useVideoList = () =>
  useAsyncData('videos', fetchVideoList, {
    default: (): Video[] => [],
    lazy: true,
  })
