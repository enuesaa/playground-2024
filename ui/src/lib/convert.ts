import type { Registry } from './registry'

export const convertToSvg = (registry: Registry): string => {
	let svg = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1000 600" width="1000" height="600">\n'

	// background
	svg += '<rect width="100%" height="100%" fill="#ffffff"/>\n'

	// shape
	for (const shape of registry.shapes) {
		if (shape.tag === 'rect') {
			svg += `<rect x="${shape.x}" y="${shape.y}" width="${shape.width}" height="${shape.height}" fill="${shape.fill}" />\n`
		} else if (shape.tag === 'text') {
			svg += `<text x="${shape.x}" y="${shape.y}">${shape.value}</text>\n`
		}
	}
	svg += '</svg>'

	return svg
}
