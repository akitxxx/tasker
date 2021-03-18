package models

import (
	"log"
	"time"
)

type Lane struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	Name      string    `json:"name"`
	IndexNum  uint64    `json:"index_num"`
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
	if err := stmt.QueryRow(id).Scan(&lane.ID, &lane.UserId, &lane.Name, &lane.IndexNum, &lane.CreatedAt, &lane.UpdatedAt); err != nil {
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
		if err := rows.Scan(&l.ID, &l.UserId, &l.Name, &l.IndexNum, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, err
		}
		lanes = append(lanes, l)
	}

	return lanes, nil
}

func CreateLane(lane *Lane) (*Lane, error) {
	var db, _ = DbConn()

	sql := "insert into lanes(user_id, name, index_num) values(?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// get current index_num
	currentIndexNum, err := getLastLaneIndexNum()
	if err != nil {
		return nil, err
	}

	// insert
	res, err := stmt.Exec(lane.UserId, lane.Name, currentIndexNum)
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

	sql := "update lanes set name = ?, index_num = ? where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// update
	_, err = stmt.Exec(l.Name, l.IndexNum, l.ID)
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

func getCurrentLaneIndexNum() (uint64, error) {
	lastIndexNum, err := getLastLaneIndexNum()
	if err != nil {
		return 0, err
	}

	if lastIndexNum == 0 {
		// return 0 if lane doesn't exist
		return lastIndexNum, nil
	}
	// return last index + 1 if lane exist
	return lastIndexNum + 1, nil
}

func getLastLaneIndexNum() (uint64, error) {
	var db, _ = DbConn()

	lane := Lane{}
	if err := db.QueryRow("select index_num from lanes order by index_num desc limit 1").Scan(&lane.IndexNum); err != nil {
		return 0, err
	}

	return lane.IndexNum, nil
}
