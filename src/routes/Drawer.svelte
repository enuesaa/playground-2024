<script lang="ts">
	let style = ''
	let border = 'none'

	$: {
		console.log(border)
		style = `border: ${border};`
	}
	
</script>

border
<input type="text" bind:value={border} />

<div {style}>
	hello
</div>
