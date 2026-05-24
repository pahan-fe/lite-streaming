<script setup lang="ts">
import { Skeleton } from '@/shared/components/ui/skeleton'
</script>

<template>
  <div class="overflow-hidden rounded-xl border border-border/60 bg-card">
    <Skeleton class="aspect-video w-full rounded-none" />
    <div class="flex items-start justify-between gap-3 p-4">
      <div class="min-w-0 flex-1 space-y-2">
        <Skeleton class="h-4 w-3/4" />
        <Skeleton class="h-3 w-2/5" />
      </div>
      <Skeleton class="size-9 shrink-0" />
    </div>
  </div>
</template>
