<script setup lang="ts">
import { useIntersectionObserver } from '@vueuse/core'

import { Button } from '@/shared/components/ui/button'

import VideoCard from './VideoCard.vue'
import VideoCardSkeleton from './VideoCardSkeleton.vue'

import type { Video } from '../schemas/video.schema'

type Props = {
  videoList: Video[]
  hasMore: boolean
  isLoadingMore: boolean
  loadMoreError: Error | null
}
defineProps<Props>()

const emit = defineEmits<{ loadMore: []; retryLoadMore: []; remove: [id: string] }>()

const sentinelRef = useTemplateRef('sentinel')

useIntersectionObserver(
  sentinelRef,
  ([entry]) => {
    if (entry?.isIntersecting) {
      emit('loadMore')
    }
  },
  { rootMargin: '200px' },
)

const LOAD_MORE_SKELETON_COUNT = 3
</script>

<template>
  <div
    v-if="videoList.length === 0"
    class="flex flex-col items-center justify-center gap-4 rounded-xl border border-dashed border-border/70 py-24 text-center"
  >
    <span class="grid size-14 place-items-center rounded-full bg-secondary text-muted-foreground">
      <svg viewBox="0 0 24 24" class="size-6 fill-current" aria-hidden="true">
        <path d="M8 5v14l11-7z" />
      </svg>
    </span>
    <div>
      <p class="font-display text-xl">No videos yet</p>
      <p class="mt-1 text-sm text-muted-foreground">Upload your first video to start streaming.</p>
    </div>
    <NuxtLink
      to="/upload"
      class="mt-2 inline-flex h-9 items-center rounded-md bg-primary px-4 text-sm font-medium text-primary-foreground transition-opacity hover:opacity-90"
    >
      Upload a video
    </NuxtLink>
  </div>

  <template v-else>
    <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-3">
      <VideoCard
        v-for="(video, i) in videoList"
        :key="video.id"
        :video="video"
        :style="{ animationDelay: `${i * 70}ms` }"
        class="animate-in duration-500 fill-mode-both fade-in slide-in-from-bottom-4"
        @remove="emit('remove', video.id)"
      />
      <template v-if="isLoadingMore">
        <VideoCardSkeleton v-for="n in LOAD_MORE_SKELETON_COUNT" :key="`skeleton-${n}`" />
      </template>
    </div>

    <div
      v-if="loadMoreError"
      class="mt-6 flex flex-col items-center gap-3 rounded-xl border border-destructive/40 bg-destructive/10 p-6 text-center"
    >
      <p class="text-sm text-destructive">Failed to load more videos.</p>
      <Button variant="outline" size="sm" @click="emit('retryLoadMore')">Retry</Button>
    </div>

    <div
      v-else-if="hasMore && !isLoadingMore"
      ref="sentinel"
      class="h-px w-full"
      aria-hidden="true"
    />
  </template>
</template>
