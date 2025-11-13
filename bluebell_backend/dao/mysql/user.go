package mysql

import (
	"bluebell_backend/model"
	"database/sql"
)

func CheckUserExist(username string) error {
	sqlStr := `select count(*) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func InsertUser(user *model.User) error {
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func GetUserByUsername(username string) (*model.User, error) {
	sqlStr := `select user_id, username, password from user where username = ?`
	user := new(model.User)
	err := db.Get(user, sqlStr, username)
	if err == sql.ErrNoRows {
		return nil, ErrorUserNotExist
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(id uint64) (*model.User, error) {
	sqlStr := `select user_id, username from user where user_id = ?`
	user := new(model.User)
	err := db.Get(user, sqlStr, id)
	if err == sql.ErrNoRows {
		return nil, ErrorUserNotExist
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
