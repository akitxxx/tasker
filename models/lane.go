package models

import "time"

type Lane struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TaskList  []Task    `json:"task_list"`
}

func FindLaneById(id uint64) (*Lane, error) {
	var db, _ = DbConn()

	sql := "select * from lanes where id = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	lane := Lane{}
	if err := stmt.QueryRow(id).Scan(&lane.UserId, &lane.Name, &lane.CreatedAt, &lane.UpdatedAt); err != nil {
		return nil, err
	}

	return &lane, nil
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
		if err := rows.Scan(&l.ID, &l.UserId, &l.Name, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, err
		}
		lanes = append(lanes, l)
	}

	return lanes, nil
}

func CreateLane(lane *Lane) (*Lane, error) {
	var db, _ = DbConn()

	sql := "insert into lanes(user_id, name, created_at, updated_at) values(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// insert
	res, err := stmt.Exec(lane.UserId, lane.Name, lane.CreatedAt, lane.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// get inserted task
	id, _ := res.LastInsertId()
	newLane, err := FindLaneById(uint64(id))

	return newLane, nil
}
