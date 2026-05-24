<script setup lang="ts">
import DeleteVideoButton from '@/features/video/components/DeleteVideoButton.vue'
import { Skeleton } from '@/shared/components/ui/skeleton'
import { formatBytes } from '~/shared/lib/formatBytes'
import { formatDate } from '~/shared/lib/formatDate'

const VideoPlayer = defineAsyncComponent(
  () => import('@/features/video/components/VideoPlayer.vue'),
)

const route = useRoute()
const videoId = route.params.id as string

const { data: video, status, error } = await useVideo(videoId)

const createdAt = computed(() => (video.value ? formatDate(video.value.createdAt, 'long') : ''))

if (error.value) {
  throw createError({
    statusCode: error.value.statusCode ?? 500,
    statusMessage: 'Video not found',
    fatal: true,
  })
}

const pageTitle = computed(() => video.value?.originalFilename ?? 'Video')
const pageDescription = computed(() =>
  video.value
    ? `Watch ${video.value.originalFilename} on lite-streaming.`
    : 'Watch video on lite-streaming.',
)

useSeoMeta({
  title: () => pageTitle.value,
  description: () => pageDescription.value,
  ogTitle: () => pageTitle.value,
  ogDescription: () => pageDescription.value,
  ogType: 'video.other',
  twitterTitle: () => pageTitle.value,
  twitterDescription: () => pageDescription.value,
})
</script>

<template>
  <div class="mx-auto max-w-4xl">
    <NuxtLink
      to="/"
      class="group mb-6 inline-flex items-center gap-1.5 text-sm text-muted-foreground transition-colors hover:text-foreground"
    >
      <span class="transition-transform duration-300 group-hover:-translate-x-1">&larr;</span>
      Library
    </NuxtLink>

    <div v-if="status === 'pending'" class="flex flex-col gap-5">
      <Skeleton class="aspect-video w-full rounded-xl" />
      <Skeleton class="h-9 w-1/2" />
      <Skeleton class="h-4 w-1/3" />
    </div>

    <div v-else-if="video" class="flex flex-col gap-6">
      <div
        class="overflow-hidden rounded-xl border border-border/60 bg-black shadow-[0_40px_120px_-40px_oklch(0.81_0.137_78/0.35)]"
      >
        <ClientOnly v-if="video.status === 'ready'">
          <VideoPlayer :src="`/api/videos/${videoId}/hls/index.m3u8`" />
          <template #fallback>
            <div class="aspect-video w-full animate-pulse bg-secondary" />
          </template>
        </ClientOnly>

        <div
          v-else
          class="flex aspect-video w-full flex-col items-center justify-center gap-3 bg-gradient-to-br from-secondary via-card to-background text-center"
        >
          <span class="flex items-center gap-2 text-primary">
            <span class="size-2 animate-pulse rounded-full bg-primary" />
            <span class="text-sm font-medium tracking-wide uppercase">Processing</span>
          </span>
          <p class="max-w-xs text-sm text-muted-foreground">
            This video is being transcoded. Check back in a moment.
          </p>
        </div>
      </div>

      <div class="flex items-start justify-between gap-4">
        <div class="min-w-0">
          <h1 class="font-display text-2xl font-medium tracking-tight md:text-3xl">
            {{ video.originalFilename }}
          </h1>
          <p class="mt-2 text-sm text-muted-foreground tabular-nums">
            {{ createdAt }} &middot; {{ formatBytes(video.size) }} &middot; {{ video.contentType }}
          </p>
        </div>
        <DeleteVideoButton :video-id="video.id" redirect-to="/" />
      </div>
    </div>
  </div>
</template>
