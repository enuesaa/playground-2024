<script lang="ts">
	import Code from './Code.svelte'

	export let out: string
</script>

<section class="mb-5">
	<h4 class="font-semibold">実行結果</h4>
	<Code language='console' code={out} />
</section>
