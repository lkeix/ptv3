package dblib

// Exec is called when inert, update, read
func Exec(query string, args []interface{}) bool {
	db := Connect()
	defer db.Close()

	tx, _ := db.Begin()
	_, err := tx.Exec(query, args...)
	success := false
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
		success = true
	}
	return success
}

// GetRow is called read
func GetRow(query string, args []interface{}) ([][]interface{}, bool) {
	db := Connect()
	defer db.Close()
	tx, _ := db.Begin()
	rows, _ := tx.Query(query, args...)
	col, err := rows.Columns()
	if err != nil {
		return nil, true
	}
	colnum := len(col)
	var res [][]interface{}
	for rows.Next() {
		elm := make([]interface{}, colnum)
		ptr := make([]interface{}, colnum)
		for i := range col {
			ptr[i] = &elm[i]
		}
		rows.Scan(ptr...)
		res = append(res, elm)
	}
	defer tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, true
	}
	return res, false
}
