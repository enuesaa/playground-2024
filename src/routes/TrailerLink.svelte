<script lang="ts">
	export let href: string
	export let title: string
</script>

<div>
	<a {href}>{title}</a>
</div>

<style lang="postcss">
	div {
		vertical-align: middle;
		width: 47%;
		@apply inline-block h-48 my-3 mx-1 sm:mx-2;
	}

	a {
		border: solid 1px rgba(0, 0, 0, 0.2);
		box-shadow: 0 1px 1px rgba(255, 255, 255, 0.2);
		overflow: hidden;
		@apply flex justify-center items-center;
		@apply w-full h-full font-zenmaru font-semibold rounded-lg;
		@apply p-3 sm:p-5 text-lg sm:text-xl;
		@apply hover:bg-grayblacker;
	}
</style>
