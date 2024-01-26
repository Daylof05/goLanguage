package sql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Update struct {
	ID        int64
	Date      string
	Fonction  string
	Arguments string
	Sortie    string
}

var db *sql.DB

func Connection() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "historydb",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func WriteUpdate(function string, argument string, sortie string) (int64, error) {
	result, err := db.Exec("INSERT INTO history (date, fonction, arguments, sortie) VALUES (?, ?, ?, ?)", time.Now(), function, argument, sortie)
	if err != nil {
		return 0, fmt.Errorf("addUpdate %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUpdate %v", err)
	}
	fmt.Println()

	return id, nil
}

func PrintUpdates() ([]Update, error) {
	var updates []Update

	rows, err := db.Query("SELECT * FROM history")
	if err != nil {
		return nil, fmt.Errorf("printUpdates: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var upd Update
		if err := rows.Scan(&upd.ID, &upd.Date, &upd.Fonction, &upd.Arguments, &upd.Sortie); err != nil {
			return nil, fmt.Errorf("printUpdates: %v", err)
		}
		updates = append(updates, upd)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("printUpdates: %v", err)
	}

	return updates, nil
}
