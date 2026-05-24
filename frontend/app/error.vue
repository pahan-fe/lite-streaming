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
  <div class="flex min-h-screen flex-col items-center justify-center gap-6 px-6 text-center">
    <p class="font-display text-[7rem] leading-none font-medium text-primary md:text-[10rem]">
      {{ error.statusCode }}
    </p>
    <div>
      <h1 class="font-display text-2xl font-medium tracking-tight md:text-3xl">
        {{ isNotFound ? 'Page not found' : 'Something went wrong' }}
      </h1>
      <p class="mx-auto mt-3 max-w-md text-sm text-muted-foreground">
        {{
          isNotFound
            ? "The page you're looking for doesn't exist or has been moved."
            : error.message
        }}
      </p>
    </div>
    <Button class="h-11 px-6" @click="handleClear">Back to library</Button>
  </div>
</template>
