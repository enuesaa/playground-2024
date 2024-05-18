<script lang="ts">
	import { treeItems, treeViewStore } from './tree'
	import Tree from './Tree.svelte'

	const { elements: { tree } } = $treeViewStore
</script>

<ul class="overflow-auto px-4 pb-4 pt-2" {...$tree}>
	<Tree {treeItems} />
</ul>
