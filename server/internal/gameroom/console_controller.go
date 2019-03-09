package gameroom

type ConsoleController struct {
	ID           int
	ConsoleID    int
	ControllerID int
}

func GetConsoleControllers() ([]ConsoleController, error) {
	rows, err := db.Query(`SELECT ID, ConsoleID, ControllerID from ConsoleControllers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	consoleControllers := make([]ConsoleController, 0)
	for rows.Next() {
		var s ConsoleController
		err := rows.Scan(&s.ID, &s.ConsoleID, &s.ControllerID)
		if err != nil {
			return nil, err
		}
		consoleControllers = append(consoleControllers, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return consoleControllers, nil
}

func CreateConsoleController(consoleID, controllerID int) error {
	_, err := db.Exec(`INSERT INTO ConsoleControllers (ConsoleID, ControllerID) VALUES (?, ?)`, consoleID, controllerID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConsoleController(ID int) error {
	_, err := db.Exec(`DELETE FROM ConsoleControllers where ID = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateConsoleController(ID, consoleID, controllerID int) error {
	_, err := db.Exec(`UPDATE ConsoleControllers SET ConsoleID = ?, ContollerID = ? WHERE ID = ?`, consoleID, controllerID, ID)
	if err != nil {
		return err
	}
	return nil
}
