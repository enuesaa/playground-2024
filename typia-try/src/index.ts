import typia, { tags } from 'typia'

interface User {
  id: string & tags.Format<'uuid'>
  name: string
  age: number & tags.ExclusiveMinimum<19> & tags.Maximum<111>
}

export const validateUser = typia.createIs<User>();
