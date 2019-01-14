package gameroom

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./gameroom.sqlite3?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
}

const (
	insertConsoleController = `INSERT INTO ConsoleControllers (ConsoleID, ControllerID) VALUES (?, ?)`
	insertController        = `INSERT INTO Controllers (Name, Image, Count) VALUES (?, ?, ?)`
	insertGame              = `INSERT INTO Games (Name, Count) VALUES (?, ?)`
	insertBarcode           = `INSERT INTO Barcodes (GameID, Barcode) VALUES (?, ?)`
	insertRental            = `INSERT INTO Rentals (BadgeID, ControllerID, GameID) VALUES (?, ?, ?)`
	insertTransaction       = `INSERT INTO Transactions (Type, BadgeID, StationID, GameID, controllerID, created) VALUES (?, ?, ?, ?, ?, ?)`
)

type Station struct {
	ID        int
	ConsoleID int
}

func GetStations() ([]Station, error) {
	rows, err := db.Query(`SELECT ID, ConsoleID from Stations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stations := make([]Station, 0)
	for rows.Next() {
		var s Station
		err := rows.Scan(&s.ID, &s.ConsoleID)
		if err != nil {
			return nil, err
		}
		stations = append(stations, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return stations, nil
}

func CreateStation(consoleID int) error {
	_, err := db.Exec(`INSERT INTO Stations (ConsoleID) VALUES (?)`, consoleID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStation(ID int) error {
	_, err := db.Exec(`DELETE FROM Stations where ID = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func CreateConsole(name, imagePath string) error {
	_, err := db.Exec(`INSERT INTO Consoles (Name, Image) VALUES (?, ?)`, name, imagePath)
	if err != nil {
		return err
	}
	return nil
}

type Transaction uint8

const (
	TranUnset Transaction = iota
	TranCheckout
	TranReturn
)

func createTransaction(tran Transaction) error {
	return nil
}
