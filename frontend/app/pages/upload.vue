<script setup lang="ts">
import { useUploadVideo } from '~/features/video/composables/useUploadVideo'
import { formatBytes } from '~/shared/lib/formatBytes'

const { upload, isPending, error } = useUploadVideo()

const file = ref<File | null>(null)

const handleFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement

  file.value = target.files?.[0] ?? null
}

const handleSubmit = async () => {
  if (!file.value) {
    return
  }

  await upload(file.value)
}
</script>

<template>
  <form class="flex flex-col gap-4" @submit.prevent="handleSubmit">
    <h1 class="text-2xl font-bold">Upload page</h1>

    <p v-if="error" class="text-destructive">{{ error.message }}</p>
    <p v-if="file">{{ file.name }} ({{ formatBytes(file.size) }})</p>

    <Label>
      Video file
      <Input type="file" accept="video/*" @change="handleFileChange" />
    </Label>

    <Button :disabled="!file || isPending" type="submit">
      {{ isPending ? 'Uploading...' : 'Upload' }}
    </Button>
  </form>
</template>
