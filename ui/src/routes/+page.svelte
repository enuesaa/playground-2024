<script lang="ts">
	import Markdown from 'svelte-exmarkdown'
	let md = '# Hello'
</script>

<textarea bind:value={md} />

<Markdown {md} />
