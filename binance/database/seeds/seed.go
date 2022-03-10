package seeds

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
)

type Seed interface {
	Run() error
}

type seed struct {
}

func NewSeed() Seed {
	return &seed{}
}

func (s seed) Run() error {
	query, err := ioutil.ReadFile("database/seeds/system_init.sql")
	if err != nil {
		return err
	}

	db, e := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_USER_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	))

	if e != nil {
		return e
	}

	if _, err := db.Exec(string(query)); err != nil {
		return err
	}

	_ = db.Close()
	return nil
}
