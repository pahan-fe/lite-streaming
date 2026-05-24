<script setup lang="ts">
import { Toaster } from '@/shared/components/ui/sonner'
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
