package gameroom

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

func UpdateStation(ID int, consoleID int) error {
	_, err := db.Exec(`UPDATE Stations SET ConsoleID = ? WHERE ID = ?`, consoleID, ID)
	if err != nil {
		return err
	}
	return nil
}
