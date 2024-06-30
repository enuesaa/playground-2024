export type Path = {
  tag: 'path'
  d: string
  stroke: string
}

export type Circle = {
  tag: 'circle'
  cx: number
  cy: number
  r: number
  fill: string
}

export type Rect = {
  tag: 'rect'
  x: number
  y: number
  width: number
  height: number
  fill: string
}

export type Shape = Path | Circle | Rect
