<script setup lang="ts">
import VideoCard from './VideoCard.vue'

import type { Video } from '../schemas/video.schema'

type Props = {
  videoList: Video[]
}

defineProps<Props>()
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

  <div v-else class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-3">
    <VideoCard
      v-for="(video, i) in videoList"
      :key="video.id"
      :video="video"
      :style="{ animationDelay: `${i * 70}ms` }"
      class="animate-in duration-500 fill-mode-both fade-in slide-in-from-bottom-4"
    />
  </div>
</template>
