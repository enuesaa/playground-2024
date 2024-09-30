export type Rect = {
	id: string
	tag: 'rect'
	x: number
	y: number
	width: number
	height: number
	fill: string
}
export type Text = {
	id: string
	tag: 'text'
	x: number
	y: number
	value: string
}
export type Line = {
	id: string
	tag: 'line'
	x: number
	y: number
	width: number
	height: number
	fill: string
}
export type Path = {
	id: string
	tag: 'path'
	d: string
}

export type Shape = Rect | Text | Path | Line
