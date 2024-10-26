export type Config = {
	title: string
	description: string
	published: string
	ignore: string[]
	variants: Variant[]
}

export type Variant = {
	name: string
	title: string
	open: string
	output: string
}
