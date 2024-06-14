import { WebContainer } from "@webcontainer/api"

const run = async (container: WebContainer, name: string, args: string[]): Promise<string> => {
  const ps = await container.spawn(name, args);
  const out = await ps.output.getReader().read()
  return out.value ?? ''
}

export const Container = () => {

  if (typeof window !== 'object') {
    return (<></>)
  }

  (async () => {
    const container = await WebContainer.boot()
    console.log(await run(container, 'node', ['--version']))
    console.log(await run(container, 'npm', ['install', '-g', 'reg-cli']))
    console.log(await run(container, 'reg-cli', ['--help']))


    await new Promise(resolve => setTimeout(resolve, 3000))

    const files = await container.fs.readdir('./')
    console.log(files)
    const files2 = await container.fs.readdir('./node_modules')
    console.log(files2)
  })()

  return (<></>)
}