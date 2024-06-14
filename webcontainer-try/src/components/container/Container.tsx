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
    const res = await fetch('nodeapp.js')
    if (res === null) {
      return
    }
    const resbody = await res.text()
    console.log(resbody)

    const ares = await fetch('a.png')
    if (res === null) {
      return
    }
    const abody = await ares.blob()
    const aarrayBuffer = await abody.arrayBuffer();
    const aunit8array = new Uint8Array(aarrayBuffer);
    const bres = await fetch('b.png')
    if (res === null) {
      return
    }
    const bbody = await bres.blob()
    const barrayBuffer = await bbody.arrayBuffer();
    const bunit8array = new Uint8Array(barrayBuffer);

    const container = await WebContainer.boot()
    container.mount({
      'nodeapp.js': {
        file: {
          contents: resbody,
        },
      },
      'aaa': {
        directory: {
          'a.png': {
            file: {
              contents: aunit8array,
            }
          },
        },
      },
      'bbb': {
        directory: {
          // 'a.png': {
          //   file: {
          //     contents: bunit8array,
          //   }
          // },
        }
      },
      'diff': {
        directory: {},
      }
    })
    console.log(await run(container, 'pwd', []))
    console.log(await run(container, 'node', ['--version']))
    console.log(await run(container, 'ls', ['-a']))
    console.log(await run(container, 'node', ['nodeapp.js', 'aaa', 'bbb', 'diff']))
    console.log(await run(container, 'ls', ['-a', 'diff']))

    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log(await run(container, 'ls', ['-a']))

    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log(await run(container, 'ls', ['-a']))

    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log(await run(container, 'ls', ['-a', 'diff']))

    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log(await run(container, 'ls', ['-a']))
    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log(await run(container, 'ls', ['-a']))
    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log(await run(container, 'ls', ['-a']))
    await new Promise(resolve => setTimeout(resolve, 3000))
    console.log(await run(container, 'ls', ['-a']))

  })()

  return (<></>)
}