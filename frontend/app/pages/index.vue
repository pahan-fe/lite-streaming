<script setup lang="ts">
import { useVideoList } from '~/features/video/composables/useVideoList'

import VideoList from '../features/video/components/VideoList.vue'
import VideoListSkeleton from '../features/video/components/VideoListSkeleton.vue'

const { data: videoList, status, error } = useVideoList()

useSeoMeta({
  title: 'Library',
  description: 'Browse and stream your uploaded videos.',
})
</script>

<template>
  <div class="flex flex-col gap-10">
    <header class="flex items-end justify-between gap-4 border-b border-border/60 pb-6">
      <div>
        <p class="mb-2 text-xs font-medium tracking-[0.2em] text-primary uppercase">Now playing</p>
        <h1 class="font-display text-4xl font-medium tracking-tight md:text-5xl">Library</h1>
        <p class="mt-3 max-w-md text-sm text-muted-foreground">
          Your uploaded videos, transcoded and ready to stream.
        </p>
      </div>
      <span
        v-if="status !== 'pending'"
        class="font-display text-4xl leading-none text-muted-foreground/30 tabular-nums"
      >
        {{ String(videoList.length).padStart(2, '0') }}
      </span>
    </header>

    <VideoListSkeleton v-if="status === 'pending'" />
    <div
      v-else-if="status === 'error'"
      class="rounded-xl border border-destructive/40 bg-destructive/10 p-6 text-sm text-destructive"
    >
      {{ error?.message }}
    </div>
    <VideoList v-else :video-list="videoList" />
  </div>
</template>
