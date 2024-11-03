import { drizzle } from 'drizzle-orm/node-postgres'
import { usersTable } from './schema'
import { eq } from 'drizzle-orm'

const db = drizzle({ 
  connection: { 
    connectionString: 'postgres://postgres:postgres@localhost:5432/postgres',
    ssl: false
  },
})

const user: typeof usersTable.$inferInsert = {
  name: 'john',
  age: 30,
  email: 'john@example.com',
}

await db.insert(usersTable).values(user)
console.log('New user created!')

const users = await db.select().from(usersTable)
console.log('Getting all users from the database: ', users)

await db
  .update(usersTable)
  .set({
    age: 31,
  })
  .where(eq(usersTable.email, user.email))
console.log('User info updated!')

// await db.delete(usersTable).where(eq(usersTable.email, user.email))
// console.log('User deleted!')
