<script setup lang="ts">
import Hls from 'hls.js'
import { onBeforeUnmount, onMounted } from 'vue'

type Props = {
  src: string
}

const props = defineProps<Props>()

const videoEl = useTemplateRef<HTMLVideoElement>('videoEl')
let hls: Hls

onMounted(() => {
  if (Hls.isSupported()) {
    if (!videoEl.value) {
      return
    }

    hls = new Hls()
    hls.loadSource(props.src)
    hls.attachMedia(videoEl.value)
  } else {
    if (videoEl.value) {
      videoEl.value.src = props.src
    }
  }
})

onBeforeUnmount(() => {
  if (Hls.isSupported() && hls) {
    hls.destroy()
  }
})
</script>

<template>
  <video ref="videoEl" controls class="block aspect-video w-full bg-black" />
</template>
