package gameroom

import "time"

type Transaction struct {
	ID           int
	Type         string
	BadgeID      int
	StationID    int
	GameID       int
	ControllerID int
	Created      time.Time
}

func GetTransactions() ([]Transaction, error) {
	rows, err := db.Query(`SELECT ID, Type, BadgeID, StationID, GameID, ControllerID, Created from Transactions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := make([]Transaction, 0)
	for rows.Next() {
		var s Transaction
		err := rows.Scan(&s.ID, &s.Type, &s.BadgeID, &s.StationID, &s.GameID, &s.ControllerID, &s.Created)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func CreateTransaction(t string, badgeID, stationID, gameID, controllerID int) error {
	_, err := db.Exec(`INSERT INTO Transactions (Type, BadgeID, StationID, GameID, ControllerID) VALUES (?, ?, ?, ?, ?)`, t, badgeID, stationID, gameID, controllerID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTransaction(ID int) error {
	_, err := db.Exec(`DELETE FROM Transactions where ID = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTransaction(ID int, t string, badgeID, stationID, gameID, controllerID int) error {
	_, err := db.Exec(`UPDATE Transactions SET Type = ?, BadgeID = ?, StationID = ?, GameID = ?, ContollerID = ? WHERE ID = ?`, t, badgeID, stationID, gameID, controllerID, ID)
	if err != nil {
		return err
	}
	return nil
}
