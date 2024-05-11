<div class="container p-3 m-auto">
	<main>
		<div class="w-48">
			<slot name="left" />
		</div>
		<div class="flex-auto">
			<slot name="right" />
		</div>
	</main>
</div>

<style lang="postcss">
	main {
		height: 85vh;
		@apply flex gap-5 mt-9 mx-auto p-3;
	}
</style>
