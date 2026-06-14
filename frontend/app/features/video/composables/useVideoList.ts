import { fetchVideoList } from '../api/videos'

import type { VideoListResponse } from '../schemas/video.schema'

const PAGE_LIMIT = 12

export const useVideoList = () => {
  const isLoadingMore = ref(false)
  const loadMoreError = ref<Error | null>(null)
  const lastLoadFailed = ref(false)

  // refresh resets accumulator (replaces data.value with page 1)
  const { data, status, error, refresh } = useAsyncData<VideoListResponse>(
    'videos',
    () => fetchVideoList({ limit: PAGE_LIMIT }),
    {
      default: () => ({ items: [], nextCursor: null }),
    },
  )

  const items = computed(() => data.value.items)
  const hasMore = computed(() => data.value.nextCursor !== null)

  const loadMore = async () => {
    const cursor = data.value.nextCursor

    if (lastLoadFailed.value || isLoadingMore.value || status.value !== 'success' || !cursor) {
      return
    }

    isLoadingMore.value = true
    try {
      const response = await fetchVideoList({ cursor, limit: PAGE_LIMIT })

      data.value = {
        items: [...data.value.items, ...response.items],
        nextCursor: response.nextCursor,
      }

      loadMoreError.value = null
      lastLoadFailed.value = false
    } catch (e) {
      loadMoreError.value = e instanceof Error ? e : new Error(String(e))
      lastLoadFailed.value = true
    } finally {
      isLoadingMore.value = false
    }
  }

  const retryLoadMore = () => {
    loadMoreError.value = null
    lastLoadFailed.value = false
    void loadMore()
  }

  const remove = (id: string) => {
    data.value = {
      ...data.value,
      items: data.value.items.filter((v) => v.id !== id),
    }
  }

  return {
    items,
    hasMore,
    status,
    error,
    refresh,
    isLoadingMore,
    loadMoreError,
    loadMore,
    retryLoadMore,
    remove,
  }
}
