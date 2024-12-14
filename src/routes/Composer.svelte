<script lang="ts">
	import ComposerTitle from './ComposerTitle.svelte'
	import { elstore, elcurrent, type El } from '$lib/el.svelte'

	let el = $state<El>({
		styles: '',
		classes: '',
		children: [],
	})

	$effect(() => {
		console.log('a')
		if (elcurrent.id !== undefined && elstore.hasOwnProperty(elcurrent.id)) {
			el = elstore[elcurrent.id]
		} else {
			el = {
				styles: '',
				classes: '',
				children: [],
			}
		}
	})
</script>

<ComposerTitle title='style' />
<textarea bind:value={el.styles}></textarea>

<ComposerTitle title='tailwind' />
<textarea bind:value={el.classes}></textarea>

<style lang="postcss">
	textarea {
		@apply block w-full outline-none bg-transparent border-zinc-600 border;
		@apply p-1;
	}
</style>
