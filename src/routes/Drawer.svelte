<script lang="ts">
	let style = $state('')
	let border = $state('none')

	$effect(() => {
		console.log(border)
		style = `border: ${border};`
	})
</script>

border
<input type="text" bind:value={border} />

<div {style}>
	hello
</div>
