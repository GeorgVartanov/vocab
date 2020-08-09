package storage

import (
	"github.com/georgvartanov/vocabProject/pkg/user/create"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func (u UserStorage) Create(user create.User) error {
	var MoscowTime, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	newUser := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  []byte(user.Password),
		Email:     user.Email,
		Created:   time.Now().In(MoscowTime),
		Changed:   time.Now().In(MoscowTime),
	}

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword(newUser.Password, cost)
	if err != nil {
		return err
	}
	newUser.Password =hash
	tx := u.DB.MustBegin()

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	_, err = tx.NamedExec(`INSERT INTO "myUser" ("firstName", "lastName", "email", "password", "created", "changed") VALUES (:firstName, :lastName, :email, :password, :created, :changed)`, &newUser)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
