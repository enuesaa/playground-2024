import { createCmd, type Prompt } from '@enuesaa/cmdgate'

const app = createCmd()

app.handle((prompt: Prompt) => {
  prompt.info('hello')
})

app.run()
