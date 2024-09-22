<script lang="ts">
	import EditorMenu from './EditorMenu.svelte'
	import Code from './Code.svelte'

	let md = '# Hello\nhey\n```ts\nexport let a = "b"\n```'
</script>

<EditorMenu />

<textarea bind:value={md} />
<Code language="md" code={md} />
