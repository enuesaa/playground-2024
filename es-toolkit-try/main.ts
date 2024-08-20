import { debounce, chunk } from 'es-toolkit'

// コールバック関数の発火を遅らせるらしい
// Node.js のイベントシステムに合わないので不思議だったが、見た感じ、あんまり凝ったことしていない
// see https://github.com/toss/es-toolkit/blob/main/src/function/debounce.ts
const debouncedLog = debounce(message => {
  console.log(message)
}, 1)

debouncedLog('Hello, world!')
// debouncedLog.cancel()

console.log('a')
console.log('a')
console.log('a')



// CakePHP の Hash 的な
const chunked = chunk([1, 2, 3, 4, 5, 6], 2)

console.log(chunked)
// Output: [[1, 2], [3, 4], [5, 6]]


