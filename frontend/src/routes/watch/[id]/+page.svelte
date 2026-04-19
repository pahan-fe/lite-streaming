<script lang="ts">
    import type { PageProps } from './$types'; 
    const { data }: PageProps = $props();
</script>

<section class="container">
    <h1>{data.video.title}</h1>
</section>

