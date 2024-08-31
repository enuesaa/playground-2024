<script lang="ts">
	type Link = {
		href: string
		text: string
		key?: string
	}
	export let links: Link[]
</script>

<ul>
	{#each links as link}
		<li>
			{#if link.key !== null}
				{link.key}:  
			{/if}
			<a href={link.href} target="_blank">{link.text}</a>
		</li>
	{/each}
</ul>

<style lang="postcss">
	ul {
		list-style-type: square;
		@apply block py-1 px-10 my-4 w-full;
	}
	li {
		letter-spacing: 0.1px;
		line-height: 1.8;
		overflow-wrap: anywhere;
		word-break: normal;
		line-break: strict;
		white-space: pre-wrap;
  		font-feature-settings: "palt";
		@apply font-zenmaru text-lg font-normal;
	}
	a {
		text-decoration: underline;
	}
</style>
