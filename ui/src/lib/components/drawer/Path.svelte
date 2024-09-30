<script lang="ts">
	import type { Path } from '$lib/shape'

	export let shape: Path
</script>

<path d={shape.d} fill="none" stroke="#000000" stroke-width="2" />
