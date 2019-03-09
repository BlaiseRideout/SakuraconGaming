package gameroom

type Game struct {
	ID        int
	Name      string
	ConsoleID int
	Count     int
}

func GetGames() ([]Game, error) {
	rows, err := db.Query(`SELECT ID, Name, ConsoleID, Count from Games`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	games := make([]Game, 0)
	for rows.Next() {
		var s Game
		err := rows.Scan(&s.ID, &s.Name, &s.ConsoleID)
		if err != nil {
			return nil, err
		}
		games = append(games, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return games, nil
}

func CreateGame(name string, consoleID int, count int) error {
	_, err := db.Exec(`INSERT INTO Games (Name, ConsoleID, Count) VALUES (?, ?, ?)`, name, consoleID, count)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGame(ID int) error {
	_, err := db.Exec(`DELETE FROM Games where ID = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateGame(ID int, name string, consoleID int, count int) error {
	_, err := db.Exec(`UPDATE Games SET Name = ?, ConsoleID = ?, Count = ? WHERE ID = ?`, name, consoleID, count, ID)
	if err != nil {
		return err
	}
	return nil
}
