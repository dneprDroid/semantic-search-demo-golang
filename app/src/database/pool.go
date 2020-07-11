package database 

import (
	"os"
	"fmt"
	"database/sql"

	"common/log"

	_ "github.com/lib/pq"
)

type Pool struct {
	db *sql.DB
}

func NewPool() *Pool {
	return &Pool{}
}

func (self *Pool) initDbConnection() error {
	db, err := sql.Open("postgres", buildConnectionString())
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		return err
	}
	self.db = db 
	return nil  
}

func (self *Pool) InContextSync(callback func(*sql.DB) error) error {
	if self.db == nil {
		if err := self.initDbConnection(); err != nil {
			return err 
		}
	}
	return callback(self.db)
}

func buildConnectionString() string {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	if user == "" || pass == "" {
		log.Fatalln("You must include POSTGRES_USER and POSTGRES_PASSWORD environment variables")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	if host == "" || port == "" || dbname == "" {
		log.Fatalln("You must include POSTGRES_HOST, POSTGRES_PORT, and POSTGRES_DB environment variables")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
}