<script lang="ts">
	import { page } from '$app/stores'
	import { normalizePath } from '$lib/path'
	import Files from './Files.svelte'
	import GoToParentDir from './GoToParentDir.svelte'

	const path = normalizePath($page.url.searchParams.get('path'))
</script>

<GoToParentDir path={path} />
<Files path={path} />
