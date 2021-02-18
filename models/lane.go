package models

import (
	"log"
	"time"
)

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
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	lane := Lane{}
	if err := stmt.QueryRow(id).Scan(&lane.ID, &lane.UserId, &lane.Name, &lane.CreatedAt, &lane.UpdatedAt); err != nil {
		log.Println(err)
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

	sql := "insert into lanes(user_id, name) values(?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// insert
	res, err := stmt.Exec(lane.UserId, lane.Name)
	if err != nil {
		return nil, err
	}

	// get inserted task
	id, _ := res.LastInsertId()
	newLane, err := FindLaneById(uint64(id))

	return newLane, nil
}

func UpdateLane(l *Lane) (*Lane, error) {
	var db, _ = DbConn()

	sql := "update lanes set name = ? where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// update
	_, err = stmt.Exec(l.Name, l.ID)
	if err != nil {
		return nil, err
	}

	// get updated task
	lane, _ := FindLaneById(l.ID)

	return lane, nil
}

func DeleteLane(id int) error {
	var db, _ = DbConn()

	stmtDelete, err := db.Prepare("delete lanes, tasks from lanes left join tasks on lanes.id = tasks.lane_id where lanes.id = ?")
	if err != nil {
		return err
	}
	defer stmtDelete.Close()

	_, err = stmtDelete.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
