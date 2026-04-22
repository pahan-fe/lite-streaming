<script lang="ts">
    import Hls from 'hls.js';
    import type { PageProps } from './$types';
    import { deleteVideoById } from '$lib/features/video/api';
    import { goto } from '$app/navigation';

    let videoEl: HTMLVideoElement | undefined = $state(); 
    const { data }: PageProps = $props();

    $effect(() => {
            if (!videoEl) {
                return;                                                    
            }
                    
            if (data.video.status === 'ready') {                                      
                const hls = new Hls();
                hls.loadSource(`/api/videos/${data.video.id}/hls/index.m3u8`);        
                hls.attachMedia(videoEl);
                                                                                    
                return () => hls.destroy();
            } else {                                                                  
                videoEl.src = `/api/videos/${data.video.id}/stream`;                  
            }
      });

        async function removeVideo() {
            if (!confirm('Delete this video?')) {
                return;
            }

            try {
                await deleteVideoById(data.video.id, fetch);
                goto('/');
            } catch (error) {
                // notification about deleting error
                console.error('Error deleting video:', error);
            }
        }
</script>

<section class="container">
    <h1>{data.video.title}</h1>
    <video bind:this={videoEl} controls></video>
    <button onclick={removeVideo}>Remove</button>
</section>

