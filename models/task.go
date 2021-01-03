package models

import (
	"time"
)

type Task struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SelectTaskList() (*[]Task, error) {
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
		if err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return &tasks, nil
}
