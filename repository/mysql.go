package repository

type Mysql struct{}

func (repo *Mysql) GetUsername() string {
	return "alice"
}
