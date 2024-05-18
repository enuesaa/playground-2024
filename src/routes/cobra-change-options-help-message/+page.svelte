<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import { type Data } from './+page'

	// see https://maku.blog/p/wqt76qw/
	export let data: Data
</script>

<PageTitle title="Cobra Change Opitons Help Message" />

<CodeViewer treeData={data.treeData} />
