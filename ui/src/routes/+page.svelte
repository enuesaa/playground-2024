<script lang="ts">
	let content = ""
</script>

<section class="w-full flex h-full">
	<div class="w-[1000px] h-[800px] px-5 py-8 border-2">
		{content}
	</div>
	<div class="flex-auto">
		<textarea bind:value={content} />
	</div>
</section>

<style lang="postcss">
	textarea {
		@apply block h-full bg-[rgba(0,0,0,0)];
		@apply font-normal w-full rounded px-3 py-2 block h-[80vh] text-black text-lg outline-none;
	}
</style>
