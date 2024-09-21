export const normalizePath = (path: string|null): string => {
  if (path === '') {
    return './'
  }
  if (path === null) {
    return './'
  }
  return path
}

export const calcParentPath = (path: string): string => {
  if (path.startsWith('./')) {
    return `.${path}`
  }
  const splitted = path.split('/')
  if (splitted.length > 0) {
    const lastItem = splitted[splitted.length - 1]
    if (lastItem !== '') {
      splitted.pop()
      return splitted.join('/') + '/'
    }
  }
  return `../${path}`
}
