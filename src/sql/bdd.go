package sql

import (
	"database/sql"
	"fmt"
	"log"

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

	fmt.Println("Connecté!!!")

	update := Update{
		Date:      "26/01/2024 11:59",
		Fonction:  "CreateFolder",
		Arguments: "NouveauDossier",
		Sortie:    "Réussie",
	}
	id, err := newUpdate(update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New update id:", id)

	updates, err := printUpdates()
	if err != nil {
		log.Fatal(err)
	}
	for _, update := range updates {
		fmt.Printf("Update found: %+v\n", update)
	}

	// updateprint, err := printUpdate()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Update found: ", updateprint)
}

func newUpdate(update Update) (int64, error) {
	result, err := db.Exec("INSERT INTO history (date, fonction, arguments, sortie) VALUES (?, ?, ?, ?)", update.Date, update.Fonction, update.Arguments, update.Sortie)
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

func printUpdates() ([]Update, error) {
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

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("printUpdates: %v", err)
	}

	return updates, nil
}

// func printUpdate() (Update, error) {
// 	var update Update

// 	row := db.QueryRow("SELECT * FROM history ")
// 	if err := row.Scan(&update.ID, &update.Date, &update.Fonction, &update.Arguments, &update.Sortie); err != nil {
// 		if err == sql.ErrNoRows {
// 			return update, fmt.Errorf("no update")
// 		} else {
// 			return update, fmt.Errorf("updates: %v", err)
// 		}
// 	}

// 	return update, nil
// }
