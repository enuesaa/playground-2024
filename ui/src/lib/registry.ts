import type { Shape } from './shape'

export type Position = {
  x: number;
  y: number;
}

export type SvgOnClick = (position: Position) => void
export type SvgOnMouseMove = (position: Position) => void
export type SvgOnMouseUp = (position: Position) => void
export type SvgOnMouseLeave = (position: Position) => void

export type Registry = {
  shapes: Shape[]
  svgOnClick: SvgOnClick|undefined
  svgOnMouseMove: SvgOnMouseMove|undefined
  svgOnMouseUp: SvgOnMouseUp|undefined
  svgOnMouseLeave: SvgOnMouseLeave|undefined
}
