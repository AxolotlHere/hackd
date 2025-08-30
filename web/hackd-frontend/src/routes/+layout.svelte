<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" class="bg-white" href={favicon} />
</svelte:head>

{@render children?.()}
