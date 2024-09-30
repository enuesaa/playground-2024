import type { Registry } from './registry'

export const createRegistry = (): Registry => {
	return {
		svg: { x: 0, y: 0 },
		shapes: [],
		histories: [],
		svgOnClick: undefined,
		svgOnMouseMove: undefined,
		svgOnMouseUp: undefined,
		svgOnMouseLeave: undefined,
		svgOnMouseDown: undefined,
		rectOnMouseDown: undefined,
	}
}
