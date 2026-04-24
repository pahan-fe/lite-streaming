<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import '../app.css';

	import { page } from '$app/state';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<header>
	<nav>
		<a href="/" class:active={page.url.pathname === '/'}>Main</a>
		<a href="/upload" class:active={page.url.pathname.startsWith('/upload')}>Upload</a>
	</nav>
</header>
<main>
	{@render children()}
</main>
