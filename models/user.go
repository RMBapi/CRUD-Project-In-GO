package models

import (
	"errors"
	"fmt"

	"example.com/project_1/db"
	"example.com/project_1/utils"
)

type User struct {
	Id       int64
	Email    string `binding: "required`
	Password string `binding: "required`
}

func (u *User) Save() error {
	query := `
   INSERT INTO users(email,password)
   VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashpassword, err := utils.Hashpassword(u.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashpassword)
	if err != nil {
		return err
	}

	User_id, err := result.LastInsertId()
	u.Id = User_id
	u.Password = hashpassword
	return err

}

func (u *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrivePassword string
	err := row.Scan(&u.Id, &retrivePassword)
	if err != nil {
		return errors.New("Wrong info")
	}
	fmt.Println(retrivePassword)

	PassValied := utils.CheckPasswordHash(u.Password, retrivePassword)
	fmt.Println(PassValied)

	if !PassValied {
		return errors.New("Wrong info")
	}

	return nil

}
