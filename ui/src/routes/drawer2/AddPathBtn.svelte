<script lang="ts">
	export let svgOnClick;

	function handleClick() {
		svgOnClick = () => {
			console.log('path')
		}
	}
</script>

<button on:click|preventDefault={handleClick}>path</button>
