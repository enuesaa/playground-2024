<script lang="ts">
	import Highlight, { HighlightAuto, LineNumbers } from 'svelte-highlight'
	import typescript from 'svelte-highlight/languages/typescript'
	import nightOwl from 'svelte-highlight/styles/night-owl'
	console.log(typescript)
</script>

<svelte:head>
	{@html nightOwl}
</svelte:head>
  
<HighlightAuto code='const a = "b"' />
<br />
<Highlight language={typescript} code='const a = "b"' let:highlighted>
  <LineNumbers {highlighted} hideBorder />
</Highlight>
