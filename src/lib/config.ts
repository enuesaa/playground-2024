export type Config = {
	project: {
		description: string,
	},
	variants: Record<string, {output: string}>
}
