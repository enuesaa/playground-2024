import { DatabaseSync } from 'node:sqlite'

// see https://nodejs.org/api/sqlite.html#sqlite
const database = new DatabaseSync(':memory:')

database.exec(`
  CREATE TABLE data(
    key INTEGER PRIMARY KEY,
    value TEXT
  ) STRICT
`);

const insert = database.prepare('INSERT INTO data (key, value) VALUES (?, ?)');
insert.run(1, 'hello');

const query = database.prepare('SELECT * FROM data');
const results = query.all()

console.log(results);

// オブジェクトにマッピングされている
console.log(results[0].key)
console.log(results[0].value)