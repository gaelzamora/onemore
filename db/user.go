package db

import (
	"database/sql"
	"fmt"
	"onemore/models"
)

func SelectUser(UserId string) (models.User, error) {
	user := models.User{}

	err := DBConnect()

	if err != nil {
		return user, err
	}
	defer Db.Close()

	query := "SELECT * FROM User WHERE Id_User = '"+UserId+"'"

	var rows *sql.Rows
	rows, err = Db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var weight sql.NullFloat64
	var height sql.NullFloat64
	var dateUpg sql.NullTime
	var age sql.NullInt64

	rows.Scan(&user.Id_User, &firstName, &lastName, &user.UserEmail, &age, &weight, &height, &user.UserStatus, &user.UserDateAdd, &dateUpg)

	user.First_Name = firstName.String
	user.Last_Name = lastName.String
	user.Weight = float32(weight.Float64)
	user.Height = float32(height.Float64)
	user.Age = int(age.Int64)
	user.UserDateUpd = dateUpg.Time.String()

	fmt.Println("Select User < Execution Successful")

	return user, nil
}