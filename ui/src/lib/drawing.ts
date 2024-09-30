import type { Registry } from './registry'
import { browser } from '$app/environment'

export const createRegistry = (): Registry => {
	return {
		svg: { x: 0, y: 0 },
		shapes: [],
		histories: [],
		svgOnClick: undefined,
		svgOnMouseMove: undefined,
		svgOnMouseUp: undefined,
		svgOnMouseLeave: undefined,
		svgOnMouseDown: undefined,
		rectOnMouseDown: undefined,
	}
}

export const createDrawing = (registry: Registry): [string, Registry] => {
	registry.histories = [structuredClone(registry.shapes)]
	const id = crypto.randomUUID()
	const data = JSON.stringify({ shapes: registry.shapes, histories: registry.histories })
	localStorage.setItem(`drawing.${id}`, data)
	return [id, registry]
}

export const updateDrawing = (id: string, registry: Registry): [string, Registry] => {
	if (registry.histories.length > 4) {
		registry.histories.shift()
	}
	registry.histories.push(structuredClone(registry.shapes))
	const data = JSON.stringify({ shapes: registry.shapes, histories: registry.histories })
	localStorage.setItem(`drawing.${id}`, data)

	return [id, registry]
}

export const deleteDrawing = (id: string) => {
	localStorage.removeItem(`drawing.${id}`)
}

export const getDrawing = (id: string): Registry => {
	const registry = createRegistry()
	if (!browser) {
		return registry
	}
	const savedata = localStorage.getItem(`drawing.${id}`)
	if (savedata !== null) {
		const data = JSON.parse(savedata)
		registry.shapes = data.shapes ?? []
		registry.histories = data.histories ?? []
		return registry
	}
	return registry
}

export const listDrawingIds = (): string[] => {
	let list: string[] = []

	if (!browser) {
		return list
	}

	for (let i = 0; i < localStorage.length; i++) {
		const key = localStorage.key(i)
		if (!key?.startsWith('drawing.')) {
			continue
		}

		const id = key.replace('drawing.', '')
		list.push(id)
	}

	return list
}
