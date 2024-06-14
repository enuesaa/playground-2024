import { css, useTheme } from '@emotion/react'

type Props = {
  name: string
}
export const Title = ({ name }: Props) => {
  const theme = useTheme()

  const styles = {
    h2: css(theme.heading, {
      width: '100%',
      color: '#cccccc',
    }),
  }

  return <h2 css={styles.h2}>{name}</h2>
}
