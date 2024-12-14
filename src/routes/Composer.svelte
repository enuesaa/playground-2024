<script lang="ts">
	import ComposerTitle from './ComposerTitle.svelte'

	let styles = $state('')
	let classes = $state('')
</script>

<ComposerTitle title='style' />
<textarea bind:value={styles}></textarea>

<ComposerTitle title='tailwind' />
<textarea bind:value={classes}></textarea>

<style lang="postcss">
	textarea {
		@apply block w-full outline-none bg-transparent border-zinc-600 border;
		@apply p-1;
	}
</style>