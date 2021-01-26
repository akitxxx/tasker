package models

import "time"

type Lane struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TaskList  []Task    `json:"task_list"`
}

func SelectLaneList() ([]Lane, error) {
	var db, _ = DbConn()

	sql := "select * from lanes;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var lanes []Lane
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		l := Lane{}
		if err := rows.Scan(&l.ID, &l.Name, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, err
		}
		lanes = append(lanes, l)
	}

	return lanes, nil
}
