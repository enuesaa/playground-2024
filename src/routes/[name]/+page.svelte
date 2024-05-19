<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import VariantTitle from './VariantTitle.svelte'
	import Provider from './Provider.svelte'
	import Readme from './Readme.svelte'
	import type { PageData } from './$types'
	import Output from './Output.svelte'

	// see https://maku.blog/p/wqt76qw/
	export let data: PageData
</script>

<PageTitle title="Change Options Help Message" />
<Readme content={data.readme} />

{#each Object.entries(data.variants) as [name, {out, files}]}
	<Provider>
		<VariantTitle title={name} />
		<CodeViewer treeData={files} />
		<Output {out} />
	</Provider>
{/each}
