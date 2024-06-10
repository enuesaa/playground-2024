import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

const user = await prisma.user.create({
  data: {
    name: 'Example',
    email: 'example@example.com',
  },
})
// { id: 1, email: 'example@example.com', name: 'Example' }
console.log(user)
await prisma.$disconnect()
