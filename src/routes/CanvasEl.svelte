<script lang="ts">
	import { createContextMenu, melt } from '@melt-ui/svelte'
	const {
	  elements: { menu, item, trigger }
	} = createContextMenu()
</script>

<div use:melt={$trigger}>
	a
</div>

<ol use:melt={$menu}>
	<li use:melt={$item}>Add</li>
	<li use:melt={$item}>Remove</li>
</ol>
  
<style lang="postcss">
	ol {
		@apply w-20 shadow rounded-lg bg-zinc-100 p-1;
	}
	li {
		@apply px-2 py-1 text-zinc-800 text-sm font-semibold;
	}
</style>
