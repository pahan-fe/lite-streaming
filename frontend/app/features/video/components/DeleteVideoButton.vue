<script setup lang="ts">
import { toast } from 'vue-sonner'

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/shared/components/ui/alert-dialog'
import { Button, buttonVariants } from '@/shared/components/ui/button'

import { useDeleteVideo } from '../composables/useDeleteVideo'

type Props = {
  videoId: string
  redirectTo?: string
}

const props = defineProps<Props>()
const emit = defineEmits<{ remove: [id: string] }>()

const { deleteById, isPending, error } = useDeleteVideo()

const handleDelete = async () => {
  try {
    await deleteById(props.videoId)
    toast.success('Video deleted')
    emit('remove', props.videoId)

    if (props.redirectTo) {
      await navigateTo(props.redirectTo, { replace: true })
    }
  } catch {
    toast.error('Failed to delete video', { description: error.value?.message })
  }
}
</script>

<template>
  <AlertDialog>
    <AlertDialogTrigger as-child>
      <Button variant="destructive" :disabled="isPending">Delete</Button>
    </AlertDialogTrigger>
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Delete this video?</AlertDialogTitle>
        <AlertDialogDescription>
          This action cannot be undone. The video and its data will be permanently removed.
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction
          :class="buttonVariants({ variant: 'destructive' })"
          :disabled="isPending"
          @click="handleDelete"
        >
          {{ isPending ? 'Deleting...' : 'Delete' }}
        </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
