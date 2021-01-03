package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SelectUserList() (*[]User, error) {
	var db, _ = DbConn()

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

func FindUserById(id uint64) (*User, error) {
	var db, _ = DbConn()

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

func FindUserByEmailAndPassword(email string, password string) (*User, error) {
	var db, _ = DbConn()

	sql := "select id, email from users where email = ? and password = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := User{}
	if err := stmt.QueryRow(email, password).Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func RegistUser(email string, password string) (*User, error) {
	var db, _ = DbConn()

	sql := "insert into users(email, password) values(?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// insert
	res, err := stmt.Exec(email, password)
	if err != nil {
		return nil, err
	}

	// get inserted user
	id, _ := res.LastInsertId()
	user, _ := FindUserById((uint64(id)))

	return user, nil
}
