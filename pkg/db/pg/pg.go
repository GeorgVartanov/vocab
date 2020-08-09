package pg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres driver
	"log"
)

// PostgresDB Postgres class struct
type PostgresDB struct {
	Host     string
	Port     int
	Password string
	User     string
	Dbname   string
	DBbPath  string
	*sqlx.DB
}

// NewPostgresDB get Postgres class struct
func NewPostgresDB(host string, port int,  user string,password string, dbname string) *PostgresDB {
	dbPath := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	pg:=PostgresDB{Host: host, Port: port, Password: password, User: user, Dbname: dbname, DBbPath: dbPath}
	if err := pg.Connect(); err != nil {
		log.Fatalln(err)
	}
	return &pg
}

// Connect connection to Postgresql
func (s *PostgresDB) Connect() error {
	db, err := sqlx.Open("postgres", s.DBbPath)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.DB = db
	return nil
}

// Disconnect close connection  to Postgresql
func (s *PostgresDB) Disconnect() error {
	if err := s.Close(); err != nil {
		return err
	}
	// fmt.Println("DB if OFF")
	return nil
}
