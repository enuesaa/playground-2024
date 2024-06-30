<script lang="ts">
	export let svgOnClick;

	function handleClick() {
		svgOnClick = () => {
			console.log('circle')
		}
	}
</script>

<button on:click|preventDefault={handleClick}>circle</button>
