package db

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func ConnectDB() (*sql.DB, error) {

	connStr := `
Server=localhost\SQLEXPRESS;Database=master;Trusted_Connection=True;
`


	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQL Server LocalDB")

	return db, nil
}
