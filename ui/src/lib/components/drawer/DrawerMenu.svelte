<script lang="ts">
	import AddRectBtn from './AddRectBtn.svelte'
	import AddTextBtn from './AddTextBtn.svelte'
	import AddPathBtn from './AddPathBtn.svelte'
	import AddLineBtn from './AddLineBtn.svelte'
	import type { Registry } from '$lib/registry'
	import DragBtn from './DragBtn.svelte'
	import ResizeBtn from './ResizeBtn.svelte'
	import DrawingDownloadBtn from './DrawingDownloadBtn.svelte'

	export let registry: Registry
</script>

<div class="flex justify-between items-center gap-2">
	<AddRectBtn bind:registry />
	<AddTextBtn bind:registry />
	<AddPathBtn bind:registry />
	<AddLineBtn bind:registry />
	<DragBtn bind:registry />
	<ResizeBtn bind:registry />

	<div class="flex-grow"></div>

	<DrawingDownloadBtn {registry} />
</div>
