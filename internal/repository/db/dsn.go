package db

func (repo *Repo) dsn() string {
	return "file:data.db?_fk=1"
}
