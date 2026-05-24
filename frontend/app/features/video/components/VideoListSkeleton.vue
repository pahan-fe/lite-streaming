<script setup lang="ts">
import VideoCardSkeleton from './VideoCardSkeleton.vue'

const PLACEHOLDER_COUNT = 6
</script>

<template>
  <div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3">
    <VideoCardSkeleton v-for="n in PLACEHOLDER_COUNT" :key="n" />
  </div>
</template>
