import fs from 'node:fs/promises'

const files = await fs.readdir('../data')
console.log(files)
