package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	LaneId    uint64    `json:"lane_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IndexNum  uint64    `json:"index_num"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SelectTaskList() ([]Task, error) {
	var db, _ = DbConn()

	sql := "select * from tasks;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tasks []Task
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t := Task{}
		if err := rows.Scan(&t.ID, &t.UserId, &t.LaneId, &t.Title, &t.Content, &t.IndexNum, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func FindTaskById(id uint64) (*Task, error) {
	var db, _ = DbConn()

	sql := "select * from tasks where id = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	task := Task{}
	if err := stmt.QueryRow(id).Scan(&task.ID, &task.UserId, &task.LaneId, &task.Title, &task.Content, &task.IndexNum, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}

	return &task, nil
}

func CreateTask(task *Task) (*Task, error) {
	var db, _ = DbConn()

	sql := "insert into tasks(user_id, lane_id, title, content, index_num) values(?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// get last index_num
	lastIndexNum, err := getLastIndexNum(task.LaneId)
	if err != nil {
		return nil, err
	}

	// insert
	res, err := stmt.Exec(task.UserId, task.LaneId, task.Title, task.Content, lastIndexNum+1)
	if err != nil {
		return nil, err
	}

	// get inserted task
	id, _ := res.LastInsertId()
	newTask, err := FindTaskById(uint64(id))

	return newTask, nil
}

func UpdateTask(t *Task) (*Task, error) {
	var db, _ = DbConn()

	sql := "update tasks set title = ?, content = ?, index_num = ? where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// update
	_, err = stmt.Exec(t.Title, t.Content, t.IndexNum, t.ID)
	if err != nil {
		return nil, err
	}

	// get updated task
	task, _ := FindTaskById(t.ID)
	return task, nil
}

func DeleteTask(id int) error {
	var db, _ = DbConn()

	stmtDelete, err := db.Prepare("delete from tasks where id = ?")
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

func getLastIndexNum(laneID uint64) (uint64, error) {
	var db, _ = DbConn()

	query := "select index_num from tasks where lane_id = ? order by index_num desc limit 1"

	task := Task{}
	err := db.QueryRow(query, laneID).Scan(&task.IndexNum)
	switch {
	case err == sql.ErrNoRows:
		// return 0 if no rows
		return 0, nil
	case err != nil:
		return 0, err
	}

	return task.IndexNum, nil
}
