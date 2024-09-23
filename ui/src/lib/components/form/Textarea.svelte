<script lang="ts">
	export let value: string = ''
	export let placeholder: string = ''
	export let readonly: boolean = false
	export let label: string = 'テキスト'
	export let handleKeyup: () => void = () => {}
</script>

<label class="w-11/12 block font-bold">
	{label}
	<textarea
		bind:value
		class="font-normal w-full rounded px-3 py-2 block h-[80vh] bg-grayer text-black text-lg outline-none border-black border bg-graywhite"
		{placeholder}
		{readonly}
		on:keyup={handleKeyup}
	></textarea>
</label>
