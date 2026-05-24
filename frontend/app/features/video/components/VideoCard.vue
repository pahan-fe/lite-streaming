<script setup lang="ts">
import { formatBytes } from '~/shared/lib/formatBytes'
import { formatDate } from '~/shared/lib/formatDate'

import DeleteVideoButton from './DeleteVideoButton.vue'

import type { Video } from '../schemas/video.schema'

type Props = {
  video: Video
}

const props = defineProps<Props>()

const createdAt = computed(() => formatDate(props.video.createdAt))

const isReady = computed(() => props.video.status === 'ready')
</script>

<template>
  <article
    class="group relative overflow-hidden rounded-xl border border-border/60 bg-card transition-all duration-300 hover:border-primary/40 hover:shadow-[0_0_0_1px_oklch(0.81_0.137_78/0.18),0_24px_60px_-24px_oklch(0.81_0.137_78/0.4)]"
  >
    <div
      class="relative aspect-video overflow-hidden bg-gradient-to-br from-secondary via-card to-background"
    >
      <div
        class="absolute inset-0 [background-image:repeating-linear-gradient(0deg,transparent,transparent_2px,oklch(0_0_0/0.4)_3px)] opacity-30"
      />

      <div class="absolute inset-0 grid place-items-center">
        <span
          class="grid size-14 place-items-center rounded-full bg-background/40 text-foreground/80 ring-1 ring-border backdrop-blur-sm transition-all duration-300 group-hover:scale-110 group-hover:bg-primary group-hover:text-primary-foreground group-hover:ring-primary"
        >
          <svg viewBox="0 0 24 24" class="size-5 translate-x-px fill-current" aria-hidden="true">
            <path d="M8 5v14l11-7z" />
          </svg>
        </span>
      </div>

      <div
        class="absolute top-3 right-3 flex items-center gap-1.5 rounded-full bg-background/70 px-2.5 py-1 text-[11px] font-medium tracking-wide backdrop-blur-sm"
      >
        <span
          class="size-1.5 rounded-full"
          :class="isReady ? 'bg-emerald-400' : 'animate-pulse bg-primary'"
        />
        {{ isReady ? 'Ready' : 'Processing' }}
      </div>
    </div>

    <div class="flex items-start justify-between gap-3 p-4">
      <div class="min-w-0">
        <h3 class="truncate font-medium tracking-tight">
          <NuxtLink
            :to="`/watch/${video.id}`"
            class="outline-none after:absolute after:inset-0 focus-visible:underline"
          >
            {{ video.originalFilename }}
          </NuxtLink>
        </h3>
        <p class="mt-1 text-xs text-muted-foreground tabular-nums">
          {{ createdAt }} &middot; {{ formatBytes(video.size) }}
        </p>
      </div>

      <span class="relative z-10 shrink-0">
        <DeleteVideoButton :video-id="video.id" />
      </span>
    </div>
  </article>
</template>
