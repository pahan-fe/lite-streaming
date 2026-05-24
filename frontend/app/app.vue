<script setup lang="ts">
import { Toaster } from '@/shared/components/ui/sonner'

useHead({
  titleTemplate: (title) => (title ? `${title} | lite-streaming` : 'lite-streaming'),
})

useSeoMeta({
  ogSiteName: 'lite-streaming',
  ogType: 'website',
  twitterCard: 'summary_large_image',
})
</script>

<template>
  <div>
    <NuxtRouteAnnouncer />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    <Toaster rich-colors position="top-right" />
  </div>
</template>
