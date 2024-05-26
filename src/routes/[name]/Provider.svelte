<script lang="ts">
	import { createTreeViewCtl, createViewing } from '$lib/tree'

	createTreeViewCtl()
	createViewing()
</script>

<div>
	<slot />
</div>

<style lang="postcss">
	div {
		border-left: solid 10px rgba(0, 0, 0, 0.2);
		@apply px-5 py-1 my-5;
	}
</style>
