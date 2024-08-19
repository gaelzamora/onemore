package db

import (
	"database/sql"
	"fmt"
	"log"
	"onemore/models"
	"onemore/secret"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err
}

func DBConnect() error {
	Db, err = sql.Open("mysql", ConvertToString(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(Db.Ping())
	err = Db.Ping()

	if err != nil {
		return err 
	}

	return nil
}

func ConvertToString(secretRDS models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = secretRDS.Username
	dbName = "gambit"
	authToken = secretRDS.Password
	dbEndpoint = secretRDS.Host
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)

	log.Println(dsn)
	return dsn
}

func UserExists(UserUUID string) (error, bool) {
	err := DBConnect()

	if err != nil {
		return err, false
	}
	defer Db.Close()

	query := "SELECT 1 FROM USER WHERE User_UUID='"+UserUUID+"'"
	log.Println(query)

	rows, err := Db.Query(query)

	if err != nil {
		return err, false
	}

	var value string
	rows.Next()
	rows.Scan(&value)

	fmt.Println("User Exists < Return Value"+value)

	if value == "1" {
		return nil, true
	}

	return nil, false
}

func IsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza IsAdmin")

	err := DBConnect()

	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	sentence := "SELECT 1 FROM User WHERE Id_User='" +userUUID+ "' AND User_Status = 0"
	fmt.Println(sentence)

	rows, err := Db.Query(sentence)

	if err != nil {
		return false, err.Error()
	}

	var value string
	rows.Next()
	rows.Scan(&value)

	fmt.Println("UserIsAdmin > Succesful - Return value " + value)
	if value == "1" {
		return true, ""
	}

	if value == "0" {
		return false, "No es admin"
	}

	return false, "User is not Admin"
}