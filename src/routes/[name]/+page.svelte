<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import type { PageData } from './$types'

	// see https://maku.blog/p/wqt76qw/
	export let data: PageData
</script>

<PageTitle title="Change Options Help Message" />

<CodeViewer treeData={data.treeData} />
