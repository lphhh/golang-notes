package migrations

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	goMigrate "github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/sirupsen/logrus"
	"os"
)

type migrate struct {
}

type Migrate interface {
	Run() error
}

func NewMigrate() Migrate {
	return &migrate{}
}

func (helper migrate) Run() error {

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_USER_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	))

	defer db.Close()

	if err != nil {
		return err
	}

	driver, e1 := mysql.WithInstance(db, &mysql.Config{})
	if e1 != nil {
		return e1
	}

	m, e2 := goMigrate.NewWithDatabaseInstance("file://database/migrations", "mysql", driver)
	if e2 != nil {
		return e2
	}

	if e3 := m.Up(); e3 != nil {
		if e3.Error() == "no change" {
			logrus.Infof("Nothing to migrate.")
			return nil
		}
		return e3
	}

	return nil
}
