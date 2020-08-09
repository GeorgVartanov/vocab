package main

import (
	"github.com/BurntSushi/toml"
	"github.com/georgvartanov/vocabProject/pkg/db/pg"
	"github.com/georgvartanov/vocabProject/pkg/user"
	"github.com/georgvartanov/vocabProject/pkg/user/create"
	"github.com/georgvartanov/vocabProject/pkg/user/login"
	"github.com/georgvartanov/vocabProject/pkg/user/read"
	"github.com/georgvartanov/vocabProject/pkg/user/storage"
	"log"
)

type TomlConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Dbname   string `toml:"dbname"`
}

func main() {

	var conf TomlConfig
	if _, err := toml.DecodeFile("config/config.toml", &conf); err != nil {
		log.Println(err)
	}
	postgresDB := pg.NewPostgresDB(conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)
	defer func() {
		_ = postgresDB.Disconnect()
	}()

	userStorage := storage.NewUserStorage(postgresDB)
	userCreate := create.NewService(userStorage)
	userRead := read.NewService(userStorage)
	userLoging := login.NewService(userStorage)
	user.Handler(userCreate, userRead, userLoging)

}
