import { globalStyle, style } from '@vanilla-extract/css'

const main = style({
  height: '50px',
  lineHeight: '50px',
  fontSize: '25px',
  fontWeight: 'bold',
})

globalStyle(`${main} svg`, {
  verticalAlign: 'middle',
  margin: '0 10px',
})

export default {
  main,
  heading: style({
    color: 'white',
    margin: '0 10px',
    textDecoration: 'none',
  }),
  setting: style({
    display: 'block',
    width: '100px',
    color: 'white',
    textDecoration: 'none',
  }),
}