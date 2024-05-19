<script lang="ts">
	// see https://github.com/metonym/svelte-highlight/issues/246
	import highlightjs from 'highlight.js/lib/common'
	import 'highlight.js/styles/night-owl.css'

	export let language: string
	export let code: string

	const hl = highlightjs.getLanguage(language)
		? highlightjs.highlight(code, { language })
		: highlightjs.highlightAuto(code)
</script>

<pre><code class="hljs">{@html hl.value}</code></pre>

<style lang="postcss">
	pre {
		@apply h-full;
	}
	code {
		@apply h-full;
	}
</style>
