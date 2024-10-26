import fs from 'node:fs/promises'

const files = await fs.readdir('../data', { withFileTypes: true })
console.log(files)
