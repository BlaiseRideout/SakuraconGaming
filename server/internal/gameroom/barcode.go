package gameroom

type Barcode struct {
	ID      int
	GameID  int
	Barcode string
}

func GetBarcodes() ([]Barcode, error) {
	rows, err := db.Query(`SELECT ID, GameID, Barcode from Barcodes`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	barcodes := make([]Barcode, 0)
	for rows.Next() {
		var s Barcode
		err := rows.Scan(&s.ID, &s.GameID, &s.Barcode)
		if err != nil {
			return nil, err
		}
		barcodes = append(barcodes, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return barcodes, nil
}

func CreateBarcode(consoleID int, barcode string) error {
	_, err := db.Exec(`INSERT INTO Barcodes (GameID, Barcode) VALUES (?, ?)`, consoleID, barcode)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBarcode(ID int) error {
	_, err := db.Exec(`DELETE FROM Barcodes where ID = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBarcode(ID int, gameID int, barcode string) error {
	_, err := db.Exec(`UPDATE Barcodes SET GameID = ?, Barcode = ? WHERE ID = ?`, gameID, ID)
	if err != nil {
		return err
	}
	return nil
}
