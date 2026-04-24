<script lang="ts">
	import { goto } from '$app/navigation';
	import { uploadVideo } from '$lib/features/video/api';
	import { formatSize } from '$lib/utils/format';

	let loading = $state<'idle' | 'loading'>('idle');
	let filesList = $state<FileList | null>(null);

	let selectedFile = $derived(filesList && filesList.length > 0 ? filesList[0] : null);

	const handleSubmit = async (e: Event) => {
		e.preventDefault();
		if (!selectedFile) {
			return;
		}

		loading = 'loading';
		try {
			const { id } = await uploadVideo(selectedFile);
			goto(`/watch/${id}`);
		} catch (err) {
			console.error('Upload error:', err);
		} finally {
			loading = 'idle';
		}
	};
</script>

<section class="space-y-20">
	<div class="space-y-10">
		<a
			href="/"
			class="group inline-flex items-center gap-2 font-mono text-[11px] uppercase tracking-[0.3em] text-ink-muted hover:text-ink transition-colors reveal"
		>
			<span class="inline-block transition-transform duration-300 group-hover:-translate-x-1">←</span>
			<span>Back to Catalogue</span>
		</a>

		<header class="space-y-7 reveal" style="animation-delay: 0.1s">
			<p class="font-mono text-[11px] uppercase tracking-[0.4em] text-amber">New Entry · Intake</p>
			<h1 class="font-display text-6xl md:text-8xl tracking-[-0.02em] text-ink leading-[0.95]">
				File a <em class="italic font-light text-amber">new reel</em>
			</h1>
			<div class="flex flex-wrap items-center gap-4 font-mono text-[11px] uppercase tracking-[0.25em] text-ink-muted pt-2">
				<span>Any video format</span>
				<span class="h-px w-20 bg-hairline draw-line" style="animation-delay: 0.4s"></span>
				<span class="text-ink-dim">Encoded to HLS on arrival</span>
			</div>
		</header>
	</div>

	<form onsubmit={handleSubmit} class="space-y-8 reveal" style="animation-delay: 0.2s">
		<label class="group block cursor-pointer">
			<input
				type="file"
				accept="video/*"
				bind:files={filesList}
				required
				class="sr-only"
				disabled={loading === 'loading'}
			/>
			<div
				class="relative border border-dashed border-hairline group-hover:border-amber-dim transition-all duration-500 px-8 md:px-16 py-20 md:py-28 text-center overflow-hidden"
			>
				<span class="absolute top-3 left-4 font-mono text-[9px] uppercase tracking-[0.4em] text-ink-dim">
					Slot · 001
				</span>
				<span class="absolute top-3 right-4 font-mono text-[9px] uppercase tracking-[0.4em] text-ink-dim">
					Intake
				</span>

				{#if selectedFile}
					<div class="space-y-4">
						<p class="font-mono text-[11px] uppercase tracking-[0.4em] text-amber">✓ Selected</p>
						<p class="font-display text-3xl md:text-4xl text-ink break-words px-4">{selectedFile.name}</p>
						<p class="font-mono text-[11px] uppercase tracking-[0.25em] text-ink-muted tabular-nums">
							{formatSize(selectedFile.size)} · {selectedFile.type || 'video/unknown'}
						</p>
						<p class="font-mono text-[10px] uppercase tracking-[0.3em] text-ink-dim pt-4">Click to replace</p>
					</div>
				{:else}
					<div class="space-y-4">
						<p class="font-mono text-[11px] uppercase tracking-[0.4em] text-ink-muted">Click to select a file</p>
						<p class="font-display text-3xl md:text-5xl italic text-ink-dim group-hover:text-ink-muted transition-colors duration-500">
							video_file.mp4
						</p>
						<p class="font-mono text-[10px] uppercase tracking-[0.3em] text-ink-dim pt-4">mp4 · mov · mkv · webm</p>
					</div>
				{/if}
			</div>
		</label>

		<div class="flex items-center justify-between gap-6 border-t border-hairline pt-8">
			<p class="font-mono text-[11px] uppercase tracking-[0.3em] text-ink-dim">
				{#if loading === 'loading'}
					<span class="cursor-blink text-amber">Filing entry</span>
				{:else if selectedFile}
					<span class="text-ink-muted">Ready to file</span>
				{:else}
					Awaiting selection
				{/if}
			</p>
			<button
				type="submit"
				disabled={loading === 'loading' || !selectedFile}
				class="font-mono text-[11px] uppercase tracking-[0.35em] text-amber border border-amber-dim px-8 py-3.5 transition-all duration-300 hover:bg-amber hover:text-bg hover:border-amber disabled:opacity-25 disabled:cursor-not-allowed disabled:hover:bg-transparent disabled:hover:text-amber"
			>
				Begin Upload →
			</button>
		</div>
	</form>
</section>
