<script lang="ts">
	import { untrack } from 'svelte'; 
	import { toast } from 'svelte-sonner';

	import type { PageProps } from './$types';
	import { formatSize } from '$lib/utils/format';
	import type { Video } from '$lib/features/video/types';
	import { fetchVideoList } from '$lib/features/video/api';
    import { pageSize } from '$lib/features/video/const';

	const { data }: PageProps = $props();

	let videos = $state<Video[]>(untrack(() => data.videos)); 
	let hasMore = $state<boolean>(untrack(() => data.videos.length >= pageSize));
	let loadingMore = $state<boolean>(false);
	let sentinel = $state<HTMLLIElement>();

	let currentPage = $state<number>(1);

	const loadMore = async () => {
		if (loadingMore || !hasMore) {
			return;
		}

		loadingMore = true;
		const nextPage = currentPage + 1;

		try {
			const newVideos = await fetchVideoList(fetch, nextPage, pageSize);
			videos = [...videos, ...newVideos];
			hasMore = newVideos.length >= pageSize;
			currentPage = nextPage;
		} catch (error) {
			toast.error('Failed to load more videos');
		} finally {
			loadingMore = false;
		}
	};

	$effect(() => {
		if (!sentinel || !hasMore) {
			return;
		}

		const observer = new IntersectionObserver(
			(entries) => {
				if (entries[0].isIntersecting) {
					loadMore();
				}
			},
			{ rootMargin: '200px' }
		);

		observer.observe(sentinel);

		return () => observer.disconnect();
	});
</script>

<section class="space-y-20">
	<header class="space-y-7 reveal">
		<p class="font-mono text-[11px] uppercase tracking-[0.4em] text-amber">Private Archive · Est. MMXXVI</p>
		<h1 class="font-display text-6xl md:text-8xl tracking-[-0.02em] text-ink leading-[0.95]">
			The <em class="italic font-light text-amber">Catalogue</em>
		</h1>
		<div class="flex flex-wrap items-center gap-4 font-mono text-[11px] uppercase tracking-[0.25em] text-ink-muted pt-2">
			<span class="tabular-nums">{String(videos.length).padStart(3, '0')} reel{videos.length === 1 ? '' : 's'} on file</span>
			<span class="h-px w-20 bg-hairline draw-line" style="animation-delay: 0.4s"></span>
			<span class="text-ink-dim">Sorted by most recent</span>
		</div>
	</header>

	{#if videos.length === 0}
		<div class="py-24 text-center space-y-5 reveal" style="animation-delay: 0.2s">
			<p class="font-mono text-[11px] uppercase tracking-[0.35em] text-ink-dim">The archive is empty</p>
			<p class="font-display text-3xl italic text-ink-muted">Nothing has been filed yet.</p>
			<a
				href="/upload"
				class="inline-block mt-6 font-mono text-[11px] uppercase tracking-[0.3em] text-amber border-b border-amber-dim hover:border-amber pb-1 transition-colors"
			>
				Begin the collection →
			</a>
		</div>
	{:else}
		<ul class="relative">
			{#each videos as video, i (video.id)}
				<li class="group reveal" style="animation-delay: {i < pageSize ? 0.15 + i * 0.05 : 0}s">
					<a
						href={`/watch/${video.id}`}
						class="relative block border-t border-hairline py-8 md:py-10 transition-[border-color,padding-left] duration-500 group-hover:border-amber-dim group-hover:pl-4"
					>
						<div class="flex items-start gap-6 md:gap-10">
							<span class="font-mono text-[11px] text-ink-dim tabular-nums pt-3 w-8 shrink-0 transition-colors duration-300 group-hover:text-amber">
								{String(i + 1).padStart(2, '0')}
							</span>
							<div class="flex-1 min-w-0 space-y-3">
								<h2 class="font-display text-2xl md:text-4xl leading-tight tracking-tight text-ink transition-colors duration-300 group-hover:text-amber truncate">
									{video.title}
								</h2>
								<div class="flex flex-wrap items-center gap-x-4 gap-y-1 font-mono text-[10px] md:text-[11px] uppercase tracking-[0.25em] text-ink-muted">
									<span>{video.updatedAt}</span>
									<span class="text-ink-dim">·</span>
									<span class="tabular-nums">{formatSize(video.size)}</span>
									<span class="text-ink-dim">·</span>
									{#if video.status === 'ready'}
										<span class="text-amber">[Ready]</span>
									{:else if video.status === 'failed'}
										<span class="text-crimson">[Processing Failed]</span>
									{:else}
										<span class="text-ink cursor-blink">[{video.status}]</span>
									{/if}
								</div>
							</div>
							<span class="font-mono text-xl md:text-2xl text-ink-dim pt-2 shrink-0 transition-all duration-500 group-hover:text-amber group-hover:translate-x-2">
								→
							</span>
						</div>
					</a>
				</li>
			{/each}
			{#if hasMore} 
				{#key videos.length}
					<li bind:this={sentinel} class="border-t border-hairline py-8 text-center font-mono text-[10px] uppercase tracking-[0.35em] text-ink-dim">                                                                                     
						{#if loadingMore}
							<span class="cursor-blink text-amber">Loading more reels</span>                                      
						{:else}
							<span>Scroll for more</span>
						{/if}                                                                                                    
					</li>
				{/key}
			{:else}
				<li class="border-t border-hairline py-8 text-center font-mono text-[10px] uppercase tracking-[0.35em] text-ink-dim">                                                                                                       
					End of archive · {videos.length} reel{videos.length === 1 ? '' : 's'} on file
				</li>
			{/if}
		</ul>
	{/if}
</section>
