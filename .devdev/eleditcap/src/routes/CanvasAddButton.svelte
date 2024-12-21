<script lang="ts">
	import type { EventHandler } from 'svelte/elements'
	import { elstore, elcurrent } from '$lib/el.svelte'
	import { nanoid } from 'nanoid'

	const handleClick: EventHandler = (e) => {
		e.preventDefault()

		if (elcurrent.id !== undefined) {
			const id = nanoid()
			elstore[id] = {
				styles: '',
				classes: 'bg-black/30 w-2/3 h-2/3',
				children: [],
			}
			elstore[elcurrent.id].children.push(id)
		}
	}
</script>

<button onclick={handleClick}>add</button>
