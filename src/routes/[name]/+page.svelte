<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import VariantTitle from './VariantTitle.svelte'
	import Provider from './Provider.svelte'
	import type { PageData } from './$types'

	// see https://maku.blog/p/wqt76qw/
	export let data: PageData
</script>

<PageTitle title="Change Options Help Message" />

{#each Object.entries(data) as [name, treeData]}
	<Provider>
		<VariantTitle title={name} />
		<CodeViewer treeData={treeData} />
	</Provider>
{/each}
