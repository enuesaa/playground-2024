<script lang="ts">
	export let href: string
	export let title: string
</script>

<a {href}>{title}</a>

<style lang="postcss">
	a {
		border: solid 1px rgba(0, 0, 0, 0.2);
		box-shadow: 0 1px 1px rgba(255, 255, 255, 0.2);
		@apply flex justify-center items-center;
		@apply w-2/5 h-40 p-5 m-5 text-xl font-zenkaku font-semibold rounded-lg;
		@apply hover:bg-grayer;
	}
</style>
