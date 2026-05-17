<script setup lang="ts">
import { useVideoList } from '~/features/video/composables/useVideoList'

import VideoList from '../features/video/components/VideoList.vue'

const { data: videoList, status, error } = await useVideoList()
</script>

<template>
  <div v-if="status === 'pending'">Loading videos...</div>
  <div v-else-if="status === 'error'">{{ error?.message }}</div>
  <VideoList v-else :video-list="videoList" />
</template>
