<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'

	// see https://maku.blog/p/wqt76qw/
	export let data
</script>

<PageTitle title="Change Options Help Message" />

<CodeViewer treeData={data.treeData} />
