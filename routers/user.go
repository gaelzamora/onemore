package routers

import (
	"encoding/json"
	"onemore/db"
)

func GetUser(body string, user string) (int, string) {
	_, find := db.UserExists(user)

	if !find {
		return 400, "User not exists ID: "+user
	}

	row, err := db.SelectUser(user)

	if err != nil {
		return 400, "Error in Select User "+err.Error()
	}

	respJSON, err := json.Marshal(row)

	if err != nil {
		return 400, "Error to attemp convert to JSON"
	}

	return 200, string(respJSON)
}