<script lang="ts">
	import VariantNavButton from './VariantNavButton.svelte'
	import { type VariantWithTreeData } from '$lib/types/tree'

	export let variants: VariantWithTreeData[]
	export let showing: string
</script>

<nav class="m-0 pl-3">
	{#each variants as variant}
		<VariantNavButton bind:showing name={variant.name} title={variant.title} />
	{/each}
</nav>

<style lang="postcss">
	nav {
		@apply flex;
		overflow: scroll;
	}
</style>
