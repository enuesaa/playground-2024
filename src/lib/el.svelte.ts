export type El = {
  styles: string
  classes: string
  children: string[]
}

let elcurrentStore = $state<undefined|string>('root')
export const elcurrent = {
  get id() {
    return elcurrentStore
  },
  update(id: string) {
    console.log(id)
    elcurrentStore = id
  }
}

export let elstore = $state<Record<string, El>>({
  'root': {
    styles: '',
    classes: 'bg-black/30 w-2/3 h-2/3',
    children: [],
  },
})
