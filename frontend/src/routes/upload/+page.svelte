<script lang="ts">
    import { goto } from '$app/navigation'
    import { uploadVideo } from '$lib/videoUpload';

    let loading = $state('idle');
    let filesList = $state <FileList | null>(null);

    const handleSubmit = async (e: Event) => {
        e.preventDefault();

        loading = 'loading';

        try {
            if (!filesList || filesList.length === 0) {
                // notification about no file selected
                return;
            }

            const { id } = await uploadVideo(filesList[0]);

            goto(`/watch/${id}`);
        } catch (err) {
            // notification about loading error
        } finally {
            loading = 'idle';
        }
    };
</script>

<h1>Upload Video</h1>
{#if loading === 'loading'}
    <p>Uploading...</p>
{:else}
    <form class="upload-form" onsubmit={handleSubmit}>
        <div class="form-group">
            <label for="file">Video File:</label>
            <input type="file" id="file" accept="video/*" bind:files={filesList} required />
        </div>
        <button type="submit">Upload</button>
    </form>
{/if}
