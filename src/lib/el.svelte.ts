export type El = {
  styles: string
  classes: string
  children: El[]
}

export const elstore = $state<El>({
  styles: '',
  classes: 'bg-black/30 w-2/3 h-2/3',
  children: [],
})
