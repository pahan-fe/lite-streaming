<script setup lang="ts">
import { toast } from 'vue-sonner'

import { useUploadVideo } from '~/features/video/composables/useUploadVideo'
import { formatBytes } from '~/shared/lib/formatBytes'

const { upload, isPending, error } = useUploadVideo()

useSeoMeta({
  title: 'Upload',
  description: 'Upload a new video for streaming.',
})

const file = ref<File | null>(null)

const handleFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement

  file.value = target.files?.[0] ?? null
}

const handleSubmit = async () => {
  if (!file.value) {
    return
  }

  try {
    await upload(file.value)
    toast.success('Video uploaded! Processing started.')
    await navigateTo('/', { replace: true })
  } catch {
    toast.error('Upload failed', { description: error.value?.message })
  }
}
</script>

<template>
  <form class="mx-auto flex max-w-xl flex-col gap-8" @submit.prevent="handleSubmit">
    <header>
      <p class="mb-2 text-xs font-medium tracking-[0.2em] text-primary uppercase">New upload</p>
      <h1 class="font-display text-4xl font-medium tracking-tight md:text-5xl">Upload a video</h1>
      <p class="mt-3 text-sm text-muted-foreground">
        We'll transcode it to HLS so it streams smoothly on any device.
      </p>
    </header>

    <label
      class="group relative flex cursor-pointer flex-col items-center justify-center gap-3 rounded-xl border-2 border-dashed border-border px-6 py-14 text-center transition-colors hover:border-primary/50 hover:bg-card"
      :class="file && 'border-primary/40 bg-card'"
    >
      <input type="file" accept="video/*" class="sr-only" @change="handleFileChange" />

      <span
        class="grid size-12 place-items-center rounded-full bg-secondary text-muted-foreground transition-colors group-hover:bg-primary/15 group-hover:text-primary"
      >
        <svg
          viewBox="0 0 24 24"
          class="size-5 fill-none stroke-current stroke-2"
          aria-hidden="true"
        >
          <path d="M12 16V4m0 0L7 9m5-5l5 5" stroke-linecap="round" stroke-linejoin="round" />
          <path d="M4 17v2a1 1 0 001 1h14a1 1 0 001-1v-2" stroke-linecap="round" />
        </svg>
      </span>

      <template v-if="file">
        <p class="max-w-full truncate font-medium">{{ file.name }}</p>
        <p class="text-xs text-muted-foreground tabular-nums">
          {{ formatBytes(file.size) }} &middot; click to replace
        </p>
      </template>
      <template v-else>
        <p class="font-medium">Click to choose a video</p>
        <p class="text-xs text-muted-foreground">MP4, MOV, WebM &mdash; any format works</p>
      </template>
    </label>

    <p
      v-if="error"
      class="rounded-lg border border-destructive/40 bg-destructive/10 px-4 py-3 text-sm text-destructive"
    >
      {{ error.message }}
    </p>

    <Button :disabled="!file || isPending" type="submit" class="h-11 text-sm">
      {{ isPending ? 'Uploading…' : 'Upload video' }}
    </Button>
  </form>
</template>
