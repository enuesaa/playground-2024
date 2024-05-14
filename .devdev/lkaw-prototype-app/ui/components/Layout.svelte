<div class="container p-3 m-auto">
	<main>
		<div class="w-48">
			<slot name="left" />
		</div>
		<div class="flex-auto overflow-hidden">
			<slot name="right" />
		</div>
	</main>
</div>

<style lang="postcss">
	main {
		min-height: 85vh;
		@apply flex gap-5 mt-6 mx-auto p-3;
	}
</style>
