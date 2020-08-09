package storage

func (u UserStorage) GetUserByEmail(email string) (User, error) {
	user :=User{}
	if err := u.Get(&user, `SELECT "id", "firstName", "lastName", "email", "password", "created", "changed" FROM "myUser" where "email"=$1`, email); err != nil {
		return user, err
	}
	return user, nil

}
