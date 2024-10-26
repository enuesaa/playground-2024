import type { Config } from '$lib/types/config'
import fs from 'node:fs/promises'
import YAML from 'yaml'

export const getConfigPath = (name: string): string => {
	return `./data/${name}/trailer.yaml`
}

export const readConfig = async (name: string): Promise<Config> => {
	const path = getConfigPath(name)
	const str = await fs.readFile(path, 'utf8')
	const config: Config = YAML.parse(str)
	return config
}
