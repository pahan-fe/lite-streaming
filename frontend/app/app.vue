<script setup lang="ts">
import { Toaster } from '@/shared/components/ui/sonner'

useHead({
  titleTemplate: (title) => (title ? `${title} | lite-streaming` : 'lite-streaming'),
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
