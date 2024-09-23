<script lang="ts">
	import FiletreeBtn from './FiletreeBtn.svelte'
	import PromptBtn from './PromptBtn.svelte'
	import SaveBtn from './SaveBtn.svelte'

	export let updateContent: (text: string) => void
	export let dirName: string
	export let content: string
</script>

<section>
	<FiletreeBtn {updateContent} />
	<PromptBtn {updateContent} />
	<SaveBtn {dirName} {content} />
</section>

<style lang="postcss">
	section {
		@apply absolute top-5 right-5;
	}
</style>
