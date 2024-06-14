import '@emotion/react'
import { CSSObject } from '@emotion/react'

declare module '@emotion/react' {
  export interface Theme {
    form: CSSObject
    button: CSSObject
    box: CSSObject
    innerbox: CSSObject
    card: CSSObject
    link: CSSObject
    linkCard: CSSObject
    heading: CSSObject

    color: {
      main: string
      highlight: string
      highlightHover: string
      contrast: string
      sub: string
      subHover: string
      danger: string
    }
    fontSize: {
      small: string
      normal: string
      large: string
      xlarge: string
      x2large: string
      x3large: string
    }
  }
}
