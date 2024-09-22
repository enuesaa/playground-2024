<script lang="ts">
	import { useCreateFile } from '$lib/api'

	const createfile = useCreateFile()

	function handleSave() {
		$createfile.mutateAsync({ path: 'aa.txt', content: 'bb' })
	}
</script>

<button>tree</button>
<button>prompt</button>
<button on:click|preventDefault={handleSave}>save</button>
<button>markdown-preview</button>
