package gameroom

type Console struct {
	ID    int
	Name  string
	Image string
}

func GetConsoles() ([]Console, error) {
	rows, err := db.Query(`SELECT ID, Name, Image from Consoles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	consoles := make([]Console, 0)
	for rows.Next() {
		var s Console
		err := rows.Scan(&s.ID, &s.Name, &s.Image)
		if err != nil {
			return nil, err
		}
		consoles = append(consoles, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return consoles, nil
}

func CreateConsole(name, imagePath string) error {
	_, err := db.Exec(`INSERT INTO Consoles (Name, Image) VALUES (?, ?)`, name, imagePath)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConsole(ID int) error {
	_, err := db.Exec(`DELETE FROM Consoles where ID = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateConsole(ID int, name, imagePath string) error {
	_, err := db.Exec(`UPDATE Consoles SET Name = ?, Image = ? WHERE ID = ?`, name, imagePath, ID)
	if err != nil {
		return err
	}
	return nil
}
