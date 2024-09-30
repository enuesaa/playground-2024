<script lang="ts">
	import { goto } from '$app/navigation'
	import { deleteDrawing } from '$lib/drawing'
	import * as AlertDialog from '$lib/components/ui/alert-dialog'

	export let id: string

	async function handleClick() {
		deleteDrawing(id)
		await goto('/')
	}
</script>

<AlertDialog.Root>
	<AlertDialog.Trigger class="text-sm border-solid border-[1px] rounded my-1 px-1">Delete</AlertDialog.Trigger>
	<AlertDialog.Content>
		<AlertDialog.Header>
			<AlertDialog.Title>Would you like to delete this drawing?</AlertDialog.Title>
		</AlertDialog.Header>
		<AlertDialog.Footer>
			<AlertDialog.Cancel class="text-sm border-solid border-[1px] rounded px-1">Cancel</AlertDialog.Cancel>
			<AlertDialog.Action on:click={handleClick} class="text-sm border-solid border-[1px] rounded px-1">
				Delete
			</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>
