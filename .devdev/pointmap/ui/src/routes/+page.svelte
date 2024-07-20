a

<!-- see http://tunani.net/index.php/2014/09/30/svg-rect/ -->

<svg viewBox="0 0 500 500">
	<rect x="100" y="20" width="40" height="40" />
	<rect x="100" y="60" width="40" height="40" fill="none" stroke="#000" stroke-width="2" />
	<rect x="100" y="100" width="40" height="40" fill="none" stroke="#000" stroke-width="2" stroke-dasharray="10,1" />
	<rect
		x="100"
		y="140"
		width="40"
		height="40"
		fill="none"
		stroke="#000"
		stroke-width="2"
		stroke-dasharray="10,5"
		stroke-linecap="round"
	/>
	<rect x="100" y="180" width="40" height="40" fill="#000" fill-opacity="0.5" />
	<rect x="100" y="220" width="40" height="40" fill="#000" rx="5" ry="5" />

	<circle cx="100" cy="150" r="20" />

	<text x="100" y="280" font-family="sans-serif" font-size="20"> Hello </text>

	<line x1="10" y1="50" x2="90" y2="50" stroke="black" stroke-width="2" />
	<polyline points="90,50 80,45 80,55" fill="none" stroke="black" stroke-width="2" />

	<path d="M10 50 L90 50 M90 50 L80 45 M90 50 L80 55" stroke="black" stroke-width="2" fill="none" />

	<defs>
		<marker id="arrowhead" markerWidth="10" markerHeight="7" refX="10" refY="3.5" orient="auto">
			<polygon points="0 0, 10 3.5, 0 7" fill="black" />
		</marker>
	</defs>
	<line x1="100" y1="330" x2="100" y2="400" stroke="#000" stroke-width="2" marker-end="url(#arrowhead)" />
</svg>

<style>
	:global(svg) {
		max-width: 400px;
		display: block;
	}
</style>
