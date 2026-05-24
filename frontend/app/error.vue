<script setup lang="ts">
import { Button } from '@/shared/components/ui/button'

import type { NuxtError } from '#app'

type Props = {
  error: NuxtError
}

const props = defineProps<Props>()

const isNotFound = computed(() => props.error.statusCode === 404)

const handleClear = () => clearError({ redirect: '/' })
</script>

<template>
  <div class="flex min-h-screen flex-col items-center justify-center gap-4 px-4 text-center">
    <h1 class="text-6xl font-bold">{{ error.statusCode }}</h1>
    <p class="text-xl font-medium">
      {{ isNotFound ? 'Page not found' : 'Something went wrong' }}
    </p>
    <p class="max-w-md text-muted-foreground">
      {{
        isNotFound ? "The page you're looking for doesn't exist or has been moved." : error.message
      }}
    </p>
    <Button @click="handleClear">Back to library</Button>
  </div>
</template>
