<script setup lang="ts">
import VideoCard from './VideoCard.vue'

import type { Video } from '../schemas/video.schema'

type Props = {
  videoList: Video[]
}

defineProps<Props>()
</script>

<template>
  <div v-if="videoList.length === 0">No videos available</div>
  <div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3">
    <VideoCard v-for="video in videoList" :key="video.id" :video="video" />
  </div>
</template>
