<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import { page } from '$app/state';
	import { ModeWatcher } from 'mode-watcher';
	import { Toaster } from '$lib/components/ui/sonner';

	let { children } = $props();
</script>

<ModeWatcher defaultMode="dark" />
<Toaster
	position="bottom-right"
	toastOptions={{
		classes: {
			toast: 'font-mono uppercase tracking-[0.15em] text-[11px]',
			title: 'font-display normal-case tracking-normal text-sm'
		}
	}}
/>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="min-h-screen flex flex-col">
	<header class="border-b border-hairline px-6 md:px-10 py-5 flex items-center justify-between">
		<a href="/" class="group flex items-baseline gap-2.5">
			<span class="font-display text-2xl tracking-tight text-ink italic">Lite</span>
			<span class="font-mono text-[10px] uppercase tracking-[0.35em] text-ink-muted group-hover:text-amber transition-colors duration-300">Stream</span>
		</a>
		<nav class="flex items-center gap-10 font-mono text-[11px] uppercase tracking-[0.28em]">
			<a
				href="/"
				class="relative py-1 transition-colors duration-300 {page.url.pathname === '/' ? 'text-amber' : 'text-ink-muted hover:text-ink'}"
			>
				{#if page.url.pathname === '/'}
					<span class="absolute -left-4 top-1/2 -translate-y-1/2 text-amber">●</span>
				{/if}
				Catalogue
			</a>
			<a
				href="/upload"
				class="relative py-1 transition-colors duration-300 {page.url.pathname.startsWith('/upload') ? 'text-amber' : 'text-ink-muted hover:text-ink'}"
			>
				{#if page.url.pathname.startsWith('/upload')}
					<span class="absolute -left-4 top-1/2 -translate-y-1/2 text-amber">●</span>
				{/if}
				New Entry
			</a>
		</nav>
	</header>

	<main class="flex-1 px-6 md:px-10 py-20">
		<div class="max-w-5xl mx-auto">
			{@render children()}
		</div>
	</main>

	<footer class="border-t border-hairline px-6 md:px-10 py-6 font-mono text-[10px] uppercase tracking-[0.35em] text-ink-dim flex items-center justify-between">
		<span>Reel · 001</span>
		<span class="hidden md:inline">Encoded to HLS on intake</span>
		<span>Private Collection</span>
	</footer>
</div>
