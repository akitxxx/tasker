package models

import (
	"time"
)

type Task struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	LaneId    uint64    `json:"lane_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
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
		if err := rows.Scan(&t.ID, &t.UserId, &t.LaneId, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt); err != nil {
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
	if err := stmt.QueryRow(id).Scan(&task.ID, &task.UserId, &task.LaneId, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}

	return &task, nil
}

func CreateTask(task *Task) (*Task, error) {
	var db, _ = DbConn()

	sql := "insert into tasks(user_id, lane_id, title, content) values(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// insert
	res, err := stmt.Exec(task.UserId, task.LaneId, task.Title, task.Content)
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

	sql := "update tasks set title = ?, content = ? where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// update
	res, err := stmt.Exec(t.Title, t.Content, t.ID)
	if err != nil {
		return nil, err
	}

	// get updated task
	id, _ := res.RowsAffected()
	task, _ := FindTaskById((uint64(id)))
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
