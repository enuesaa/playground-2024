import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

const users = await prisma.user.findMany({})
console.log(users)
await prisma.$disconnect()
