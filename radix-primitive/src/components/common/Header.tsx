import Link from 'next/link'
import { AiTwotoneSetting } from 'react-icons/ai'
import { FaGuitar } from 'react-icons/fa'
import { Flex, Box, Container } from '@radix-ui/themes'
import styles from '@/components/common/Header.css'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4'>
        <Flex>
          <Box grow='1'>
            <Link href='/' className={styles.heading}>
              <FaGuitar />
              my-nextjs-template
            </Link>
          </Box>

          <Link href='/setting' className={styles.setting}>
            <AiTwotoneSetting />
          </Link>
        </Flex>
      </Container>
    </header>
  )
}
