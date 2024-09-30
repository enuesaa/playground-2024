<script lang="ts">
	import { type Rect } from '$lib/shape'

	export let shape: Rect
</script>

<rect x={shape.x} y={shape.y} width={shape.width} height={shape.height} fill={shape.fill} />
