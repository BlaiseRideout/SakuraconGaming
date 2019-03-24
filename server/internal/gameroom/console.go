package gameroom

import (
	"database/sql"
)

type Console struct {
	ID          int64
	Name        string
	Image       string
	Controllers []Controller
}

func GetConsoles() ([]Console, error) {
	rows, err := db.Query(
		`SELECT Consoles.ID, Consoles.Name, Consoles.Image, NULL,           NULL             FROM Consoles UNION
		 SELECT Consoles.ID, Consoles.Name, Consoles.Image, Controllers.ID, Controllers.Name FROM Consoles
		 INNER JOIN ConsoleControllers ON Consoles.ID = ConsoleControllers.ConsoleID
		 INNER JOIN Controllers ON Controllers.ID = ConsoleControllers.ControllerID;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	consoles := make(map[int64]Console, 0)
	for rows.Next() {
		var console Console
		var controllerID sql.NullInt64
		var controllerName sql.NullString
		err := rows.Scan(&console.ID, &console.Name, &console.Image, &controllerID, &controllerName)
		if err != nil {
			return nil, err
		}
		if controllerID.Valid && controllerName.Valid {
			console.Controllers = make([]Controller, 1)
			console.Controllers[0].ID = controllerID.Int64
			console.Controllers[0].Name = controllerName.String
			existingConsole, exists := consoles[console.ID]
			if exists {
				console.Controllers = append(existingConsole.Controllers, console.Controllers...)
			}
		}
		consoles[console.ID] = console
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	consolesArr := make([]Console, len(consoles))
	idx := 0
	for _, value := range consoles {
		consolesArr[idx] = value
		idx++
	}
	return consolesArr, nil
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
