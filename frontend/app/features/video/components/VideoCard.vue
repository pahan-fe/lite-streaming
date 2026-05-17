<script setup lang="ts">
import { Badge } from '@/shared/components/ui/badge'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/shared/components/ui/card'

import type { Video } from '../schemas/video.schema'

type Props = {
  video: Video
}

const props = defineProps<Props>()

const createdAt = useDateFormat(() => props.video.createdAt, 'MMM D, YYYY')
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>{{ video.originalFilename }}</CardTitle>
      <CardDescription>{{ createdAt }}</CardDescription>
    </CardHeader>
    <CardContent>
      <Badge :variant="video.status === 'ready' ? 'default' : 'secondary'">
        {{ video.status }}
      </Badge>
    </CardContent>
  </Card>
</template>
