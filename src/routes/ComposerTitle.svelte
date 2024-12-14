<script lang="ts">
	type Props = {
		title: string
	}
	let { title }: Props = $props()
</script>

<div class="w-full py-1 px-1">
	{title}
</div>
