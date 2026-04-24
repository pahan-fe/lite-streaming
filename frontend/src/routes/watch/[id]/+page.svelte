<script lang="ts">
	import Hls from 'hls.js';
	import type { PageProps } from './$types';
	import { deleteVideoById } from '$lib/features/video/api';
	import { goto, invalidate } from '$app/navigation';
	import { toast } from 'svelte-sonner';
	import { formatSize } from '$lib/utils/format';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';

	const { data }: PageProps = $props();

	let videoEl: HTMLVideoElement | undefined = $state();
	let dialogOpen = $state(false);
	let deleting = $state(false);

	const videoId = $derived(data.video.id);
	const videoStatus = $derived(data.video.status);

	$effect(() => {
		if (!videoEl) {
			return;
		}

		if (videoStatus === 'ready') {
			const hls = new Hls();
			hls.loadSource(`/api/videos/${videoId}/hls/index.m3u8`);
			hls.attachMedia(videoEl);

			return () => hls.destroy();
		} else {
			videoEl.src = `/api/videos/${videoId}/stream`;
		}
	});

	$effect(() => {
		if (videoStatus === 'ready') {
			return;
		}

		const interval = setInterval(() => {
			invalidate(`video:${videoId}`);
		}, 5000);

		return () => clearInterval(interval);
	});

	async function confirmRemove() {
		deleting = true;
		try {
			await deleteVideoById(data.video.id, fetch);
			toast.success('Reel removed from archive');
			dialogOpen = false;
			goto('/');
		} catch (error) {
			const message = error instanceof Error ? error.message : 'Unknown error';
			toast.error(`Failed to remove · ${message}`);
			deleting = false;
		}
	}
</script>

<section class="space-y-12">
	<header class="space-y-6 reveal">
		<div class="flex items-center gap-4 font-mono text-[10px] uppercase tracking-[0.4em]">
			<span class="text-amber">Now Showing</span>
			<span class="h-px w-16 bg-amber-dim draw-line" style="animation-delay: 0.3s"></span>
			<span class="text-ink-dim tabular-nums">{data.video.id.slice(0, 8)}</span>
		</div>
		<h1 class="font-display text-5xl md:text-7xl tracking-[-0.02em] text-ink leading-[1]">
			{data.video.title}
		</h1>
		<div class="flex flex-wrap items-center gap-x-4 gap-y-1 font-mono text-[11px] uppercase tracking-[0.25em] text-ink-muted pt-2">
			<span>{data.video.updatedAt}</span>
			<span class="text-ink-dim">·</span>
			<span class="tabular-nums">{formatSize(data.video.size)}</span>
			<span class="text-ink-dim">·</span>
			{#if data.video.status === 'ready'}
				<span class="text-amber">[HLS · Ready]</span>
			{:else}
				<span class="text-ink cursor-blink">[{data.video.status}]</span>
			{/if}
		</div>
	</header>

	<div class="reveal" style="animation-delay: 0.2s">
		<div class="relative border border-hairline bg-black shadow-[0_40px_80px_-20px_oklch(0_0_0/0.5)]">
			<video bind:this={videoEl} controls class="w-full aspect-video block"></video>
			<span class="absolute top-3 left-4 font-mono text-[9px] uppercase tracking-[0.4em] text-amber/70 pointer-events-none">
				Reel · {data.video.id.slice(0, 6)}
			</span>
			<span class="absolute top-3 right-4 font-mono text-[9px] uppercase tracking-[0.4em] text-amber/70 pointer-events-none">
				Now Playing
			</span>
			<span class="absolute bottom-3 left-4 font-mono text-[9px] uppercase tracking-[0.4em] text-amber/70 pointer-events-none">
				{data.video.status === 'ready' ? 'HLS · Adaptive' : 'Raw · Pass-through'}
			</span>
			<span class="absolute bottom-3 right-4 font-mono text-[9px] uppercase tracking-[0.4em] text-amber/70 pointer-events-none tabular-nums">
				24 fps
			</span>
		</div>
	</div>

	<footer class="flex flex-wrap items-center justify-between gap-6 border-t border-hairline pt-8 reveal" style="animation-delay: 0.35s">
		<a
			href="/"
			class="group font-mono text-[11px] uppercase tracking-[0.3em] text-ink-muted hover:text-ink transition-colors"
		>
			<span class="inline-block transition-transform duration-300 group-hover:-translate-x-1">←</span>
			<span class="ml-2">Back to Catalogue</span>
		</a>
		<button
			onclick={() => dialogOpen = true}
			class="group font-mono text-[11px] uppercase tracking-[0.3em] text-ink-dim hover:text-crimson transition-colors"
		>
			<span class="border-b border-transparent group-hover:border-crimson pb-1 transition-colors">
				Remove from Archive
			</span>
		</button>
	</footer>
</section>

<Dialog.Root bind:open={dialogOpen}>
	<Dialog.Content class="bg-card border-hairline max-w-md">
		<Dialog.Header class="space-y-4">
			<p class="font-mono text-[11px] uppercase tracking-[0.4em] text-crimson">
				Confirm · Removal
			</p>
			<Dialog.Title class="font-display text-3xl md:text-4xl tracking-tight text-ink leading-tight">
				Remove this reel?
			</Dialog.Title>
			<Dialog.Description class="font-display text-base text-ink-muted leading-relaxed">
				"{data.video.title}" will be erased from the archive. The action cannot be undone.
			</Dialog.Description>
		</Dialog.Header>
		<Dialog.Footer class="flex flex-row items-center justify-end gap-4 border-t border-hairline pt-6 mt-4">
			<Button
				variant="ghost"
				onclick={() => dialogOpen = false}
				disabled={deleting}
				class="font-mono text-[11px] uppercase tracking-[0.3em]"
			>
				Keep
			</Button>
			<Button
				variant="destructive"
				onclick={confirmRemove}
				disabled={deleting}
				class="font-mono text-[11px] uppercase tracking-[0.3em]"
			>
				{deleting ? 'Removing…' : 'Remove →'}
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
