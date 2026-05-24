<script setup lang="ts">
const VideoPlayer = defineAsyncComponent(
  () => import('@/features/video/components/VideoPlayer.vue'),
)

const route = useRoute()
const videoId = route.params.id as string

const { data: video, status, error } = await useVideo(videoId)
</script>

<template>
  <div>
    <div v-if="status === 'pending'">Loading...</div>
    <div v-else-if="status === 'error'">{{ error?.message }}</div>
    <div v-else-if="video">
      <h1 class="mt-2 text-2xl font-bold">{{ video.originalFilename }}</h1>
      <ClientOnly v-if="video.status === 'ready'">
        <VideoPlayer :src="`/api/videos/${videoId}/hls/index.m3u8`" />
      </ClientOnly>
      <div v-else>Video is processing</div>
    </div>
  </div>
</template>
