<script lang="ts">
    const { data } = $props();
</script>

<section class="container">
    <h1>Video List</h1>
    <ul class="video-list">
        {#each data.videos as video}
            <li class="video-item">
                <a href={`/watch/${video.id}`}> 
                    <h2>{video.title}</h2>
                </a>
            </li>
        {/each}
    </ul>
</section>
