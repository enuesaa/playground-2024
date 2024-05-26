<script lang="ts">
	import { AlertTriangleIcon } from 'svelte-feather-icons'
</script>

<section class="container mx-auto">
	<span><AlertTriangleIcon size='20' /></span> This website is currentlty under development.
</section>

<style lang="postcss">
	section {
		background: rgba(255, 224, 14, 0.3);
		border: dashed 1.5px;
		@apply rounded-lg px-5 py-2 border-black select-none mt-1 text-center font-medium;
	}
	span {
		display: inline-block;
		vertical-align: text-bottom;
		@apply text-sm mr-1;
	}
</style>
