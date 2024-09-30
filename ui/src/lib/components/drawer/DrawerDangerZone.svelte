<script lang="ts">
	import DrawingDeleteBtn from './DrawingDeleteBtn.svelte'

	export let id: string | undefined = undefined
</script>

<div class="flex justify-end">
	{#if id !== undefined}
		<DrawingDeleteBtn {id} />
	{/if}
</div>
