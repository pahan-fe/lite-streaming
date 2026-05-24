<script setup lang="ts">
import { Badge } from '@/shared/components/ui/badge'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/shared/components/ui/card'

import DeleteVideoButton from './DeleteVideoButton.vue'

import type { Video } from '../schemas/video.schema'

type Props = {
  video: Video
}

const props = defineProps<Props>()

const createdAt = useDateFormat(() => props.video.createdAt, 'MMM D, YYYY')
</script>

<template>
  <Card class="relative transition-shadow hover:shadow-md">
    <CardHeader>
      <CardTitle>
        <NuxtLink :to="`/watch/${video.id}`" class="after:absolute after:inset-0">
          {{ video.originalFilename }}
        </NuxtLink>
      </CardTitle>
      <CardDescription>{{ createdAt }}</CardDescription>
    </CardHeader>
    <CardContent class="flex items-center justify-between">
      <Badge :variant="video.status === 'ready' ? 'default' : 'secondary'">
        {{ video.status }}
      </Badge>
      <span class="relative z-10">
        <DeleteVideoButton :video-id="video.id" />
      </span>
    </CardContent>
  </Card>
</template>
