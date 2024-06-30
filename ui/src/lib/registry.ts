export type SvgOnClick = (x: number, y: number) => void
export type SvgOnMouseMove = (x: number, y: number) => void
export type SvgOnMouseUp = (x: number, y: number) => void
export type SvgOnMouseLeave = (x: number, y: number) => void

export type Registry = {
  svgOnClick: SvgOnClick|undefined
  svgOnMouseMove: SvgOnMouseMove|undefined
  svgOnMouseUp: SvgOnMouseUp|undefined
  svgOnMouseLeave: SvgOnMouseLeave|undefined
}
