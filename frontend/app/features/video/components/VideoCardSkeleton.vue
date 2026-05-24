<script setup lang="ts">
import { Card, CardContent, CardHeader } from '@/shared/components/ui/card'
import { Skeleton } from '@/shared/components/ui/skeleton'
</script>

<template>
  <Card>
    <CardHeader class="gap-2">
      <Skeleton class="h-5 w-3/4" />
      <Skeleton class="h-4 w-1/3" />
    </CardHeader>
    <CardContent class="flex items-center justify-between">
      <Skeleton class="h-6 w-16" />
      <Skeleton class="h-9 w-20" />
    </CardContent>
  </Card>
</template>
