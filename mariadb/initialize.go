package mariadb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Maria struct {
	Worker *sql.DB
}

func NewDatabase() (m *Maria, err error) {
	user, password, dburl, dbname := os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DBURL"), os.Getenv("MYSQL_DBNAME")
	dataSourceName := user + ":" + password + "@tcp(" + dburl + ")/" + dbname

	db, err := sql.Open("mysql", dataSourceName)

	m = &Maria{
		Worker: db,
	}

	return
}

func (m *Maria) CloseDatabase() {
	m.Worker.Close()
}
