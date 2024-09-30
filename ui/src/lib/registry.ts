import type { Shape } from './shape'

export type Position = {
	x: number
	y: number
}

export type SvgOnClick = (position: Position) => void
export type SvgOnMouseMove = (position: Position) => void
export type SvgOnMouseUp = (position: Position) => void
export type SvgOnMouseLeave = (position: Position) => void
export type SvgOnMouseDown = (position: Position) => void
export type RectOnMouseDown = (shapeId: string, position: Position) => void

export type Registry = {
	svg: Position
	shapes: Shape[]
	histories: Shape[][]
	svgOnClick: SvgOnClick | undefined
	svgOnMouseMove: SvgOnMouseMove | undefined
	svgOnMouseUp: SvgOnMouseUp | undefined
	svgOnMouseLeave: SvgOnMouseLeave | undefined
	svgOnMouseDown: SvgOnMouseDown | undefined
	rectOnMouseDown: RectOnMouseDown | undefined
}
