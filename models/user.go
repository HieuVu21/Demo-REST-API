package models

import (
	"REST_API/db"
	"REST_API/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "insert into users(email, password) values (?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)	
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	UserId, err := result.LastInsertId()
	u.ID = UserId
	return nil

}
func (u *User) ValidateCredentials() error {
	query := "select id, password from users where email =?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID,&retrievedPassword)
 if err != nil{ 
	return err
 }

    passwordIsValid := utils.CheckHashPassword(u.Password,retrievedPassword)
	if !passwordIsValid{
		return errors.New("credential invalid")
	}
	return nil
}

