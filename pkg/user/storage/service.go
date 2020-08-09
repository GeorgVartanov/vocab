package storage

import (
	"github.com/georgvartanov/vocabProject/pkg/db/pg"
)

type UserStorage struct {
	*pg.PostgresDB
}



func NewUserStorage(pg *pg.PostgresDB) *UserStorage {
	return &UserStorage{PostgresDB: pg}
}




