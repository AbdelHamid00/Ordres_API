package initializer

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"fmt"
)

func ConnectDataBase() (*sql.DB, error) {
	// Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "BadrCham",
    }
    // Get a database handle.
    var err error
    DB, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return nil, err
    }

    pingErr := DB.Ping()
    if pingErr != nil {
        return nil, pingErr
    }
    fmt.Println("Connected!")
	return DB, nil
}