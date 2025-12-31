package db

import (
	"fmt"
	"project/app/db/conn"
	"project/app/model"
)

type User interface {
	Register(user model.User) conn.ExecResult
	GetUser(username string) (model.User, error)
	IsUserExist(user model.User) (string, bool)
}

type user struct{}

func NewUser() User {
	return &user{}
}

func (u *user) Register(user model.User) conn.ExecResult {
	execRes := conn.NewSqlDb().ExecSql("insert into user_info (username, password, name , email, age) values ($1, $2, $3, $4, $5)", user.Username, user.Password, user.Name, user.Email, user.Age)

	if execRes.Error != nil {
		_ = fmt.Errorf("error in query: %v", execRes.Error)
	}
	return execRes
}

func (u *user) GetUser(username string) (model.User, error) {
	fmt.Println("Username: ", username)
	queryRes := conn.NewSqlDb().QuerySql("select username, password, name, email, age from user_info where username= $1", username)
	var user model.User
	if queryRes.Error != nil {
		fmt.Printf("error in GetUser: %v", queryRes.Error)
		return user, queryRes.Error
	}
	rows := queryRes.Rows
	if rows != nil {
		for rows.Next() {
			rows.Scan(&user.Username, &user.Password, &user.Name, &user.Email, &user.Age)
		}
		rows.Close()
	}
	return user, nil
}

func (u *user) IsUserExist(user model.User) (string, bool) {
	var message string
	var res bool = false
	queryRes := conn.NewSqlDb().QuerySql("select username from user_info where username= $1", user.Username)
	var username string
	if queryRes.Error == nil {
		rows := queryRes.Rows
		if rows != nil {
			for rows.Next() {
				rows.Scan(&username)
			}
		}
	}
	if username != "" {
		message = "username exist choose different one"
		res = true
	}
	queryRes = conn.NewSqlDb().QuerySql("select email from user_info where email= $1", user.Email)
	var email string
	if queryRes.Error == nil {
		rows := queryRes.Rows
		if rows != nil {
			for rows.Next() {
				rows.Scan(&email)
			}
		}
	}
	if username != "" {
		if message != "" {
			message += " and "
		}
		message += "email ID exist"
		res = true
	}
	return message, res
}
