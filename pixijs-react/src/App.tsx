import { BlurFilter } from 'pixi.js'
import { Stage, Container, Sprite, Text } from '@pixi/react'
import { useMemo } from 'react'

// see https://pixijs.io/pixi-react/
export const App = () => {
  const blurFilter = useMemo(() => new BlurFilter(4), [])

  return (
    <Stage>
      <Sprite
        image="https://pixijs.io/pixi-react/img/bunny.png"
        x={400}
        y={270}
        anchor={{ x: 0.5, y: 0.5 }}
        filters={[blurFilter]}
      />

      <Container x={400} y={330}>
        {/* <Text text="Hello World" anchor={{ x: 0.5, y: 0.5 }} filters={[blurFilter]} /> */}
      </Container>
    </Stage>
  )
}
