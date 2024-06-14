import { css } from '@emotion/react'
import normalize from 'emotion-normalize'

export const globalStyle = css`
  ${normalize}

  html,
  body {
    padding: 0;
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, sans-serif;
  }
  a {
    color: inherit;
    text-decoration: none;
  }
  * {
    box-sizing: border-box;
  }
  body {
    background: #1a1a1a;
  }
`
