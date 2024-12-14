<script lang="ts">
	import Canvas from './Canvas.svelte'
	import Composer from './Composer.svelte'
</script>

<div class="flex gap-1 h-screen">
	<div class="w-1/2">
		<Canvas />
	</div>
	<div class="w-1/2">
		<Composer />
	</div>
</div>
