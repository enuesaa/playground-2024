export type Variant = {
	name: string
	title: string
	output: string
}

export type Config = {
	title: string
	description: string
	variants: Variant[]
}
