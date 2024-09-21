<script lang="ts">
	import TrailerLink from './TrailerLink.svelte'
	import type { PageData } from './$types'

	export let data: PageData
</script>

<svelte:head>
	<title>雑コード置き場</title>
</svelte:head>

<div class="text-center">
	{#each data.projects as project}
		<TrailerLink href={`/${project.name}`} title={project.title} />
	{/each}
</div>
