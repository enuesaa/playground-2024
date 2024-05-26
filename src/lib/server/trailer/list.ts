import fs from 'node:fs/promises'

export const listTrailers = async (): Promise<string[]> => {
	const files = await fs.readdir('./data', { withFileTypes: true })
	return files.map((f) => f.name)
}
