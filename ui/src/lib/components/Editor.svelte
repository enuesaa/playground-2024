<script lang="ts">
	import EditorMenu from './EditorMenu.svelte'

	import Markdown from 'svelte-exmarkdown'
	let md = '# Hello'
</script>

<EditorMenu />

<textarea bind:value={md} />
<Markdown {md} />
