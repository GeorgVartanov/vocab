package storage

import (
	"github.com/georgvartanov/vocabProject/pkg/user/read"
)

func (u UserStorage) Read(id int) (read.User, error) {
	user := read.User{}
	if err := u.Get(&user, `SELECT "id", "firstName", "lastName", "email", "created", "changed" FROM "myUser" where "id"=$1` , id); err != nil {
		return user, err
	}
	return user, nil

}

func (u UserStorage) ReadAll() ([]read.User, error) {
	user := []read.User{}
	if err := u.Select(&user, `SELECT "id", "firstName", "lastName", "email", "created", "changed" FROM "myUser"`); err != nil {
		return nil, err
	}
	return user, nil
}
