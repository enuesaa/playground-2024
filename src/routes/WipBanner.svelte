<script lang="ts">
	import { AlertTriangleIcon } from 'svelte-feather-icons'
</script>

<section>
	<span><AlertTriangleIcon size='18' /></span> This website is currentlty under development.
</section>

<style lang="postcss">
	section {
		background: rgba(255, 224, 14, 0.3);
		border: dashed 1.5px;
		@apply rounded-lg p-2 border-black select-none mt-1 text-center font-medium;
		@apply container mx-auto w-11/12;
	}
	span {
		display: inline-block;
		vertical-align: text-bottom;
		@apply text-sm mr-1;
	}
</style>
