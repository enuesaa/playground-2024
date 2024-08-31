<script lang="ts">
	import Description from '../[name]/Description.svelte'
	import PageTitle from '../[name]/PageTitle.svelte'
	import Links from '../[name]/Links.svelte'

	let about = 'このサイトは個人のメモ帳です。\n'
	about += 'ちょっと調べて気づいたことをメモしてます。\n\n'

	let links = [{href: 'https://enuesaa.dev/', text: 'nua', key: '運営者'}]
</script>

<PageTitle title='About' />
<Description content={about} />

<Links {links} />
