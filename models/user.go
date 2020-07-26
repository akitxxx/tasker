package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SelectUserList() (*[]User, error) {
	var db, _ = sql.Open("mysql", "root:root@tcp(tasker_dev_db)/tasker_dev")

	sql := "select id, email from users;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []User
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return &users, nil
}

func FindById(id uint64) (*User, error) {
	var db, _ = sql.Open("mysql", "root:root@tcp(tasker_dev_db)/tasker_dev")

	sql := "select id, email from users where id = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := User{}
	if err := stmt.QueryRow(id).Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
