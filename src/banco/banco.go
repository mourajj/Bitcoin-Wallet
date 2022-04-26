package banco

import (
	"database/sql"
	"maxxer/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conectar starts the connection with the database and returns it
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.DBConnectionString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}
