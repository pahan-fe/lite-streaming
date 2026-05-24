<script setup lang="ts">
import { useVideoList } from '~/features/video/composables/useVideoList'

import VideoList from '../features/video/components/VideoList.vue'
import VideoListSkeleton from '../features/video/components/VideoListSkeleton.vue'

const { data: videoList, status, error } = useVideoList()
</script>

<template>
  <VideoListSkeleton v-if="status === 'pending'" />
  <div v-else-if="status === 'error'" class="text-destructive">{{ error?.message }}</div>
  <VideoList v-else :video-list="videoList" />
</template>
