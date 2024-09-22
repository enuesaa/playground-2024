<script lang="ts">
	import { calcParentPath } from '$lib/path'
	import { ArrowUpIcon } from 'svelte-feather-icons'

	export let path: string

	const parentPath = calcParentPath(path)
</script>

<div class="flex my-5">
	<div class="p-3 pb-1 flex-initial w-12 text-center">
		<a href={`/?path=` + parentPath} class="inline-block" data-sveltekit-reload>
			<ArrowUpIcon />
		</a>
	</div>
	<div class="p-3 pb-1 flex-none bg-slate-600" style="min-width: 300px;">
		{path}
	</div>
</div>
