package gameroom

type Controller struct {
	ID    int64
	Name  string
	Image string
	Count int
}

func GetControllers() ([]Controller, error) {
	rows, err := db.Query(`SELECT ID, Name, Image, Count from Controllers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	controllers := make([]Controller, 0)
	for rows.Next() {
		var s Controller
		err := rows.Scan(&s.ID, &s.Name, &s.Image, &s.Count)
		if err != nil {
			return nil, err
		}
		controllers = append(controllers, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return controllers, nil
}

func CreateController(name, imagePath string, count int) error {
	_, err := db.Exec(`INSERT INTO Controllers (Name, Image, Count) VALUES (?, ?, ?)`, name, imagePath, count)
	if err != nil {
		return err
	}
	return nil
}

func DeleteController(ID int) error {
	_, err := db.Exec(`DELETE FROM Controllers where ID = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateController(ID int, name, imagePath string, count int) error {
	_, err := db.Exec(`UPDATE Controllers SET Name = ?, Image = ?, Count = ? WHERE ID = ?`, name, imagePath, count, ID)
	if err != nil {
		return err
	}
	return nil
}
