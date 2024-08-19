package routers

import (
	"encoding/json"
	"onemore/db"
	"onemore/models"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func InsertMuscle(body string, user string) (int, string) {
	var m models.Muscle

	err := json.Unmarshal([]byte(body), &m)

	if err != nil {
		return 400, "Error to attemp convert to JSON"+err.Error()
	}

	if len(m.Name_Muscle) == 0 {
		return 400, "Name is too short"
	}

	isAdmin, msg := db.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	result, err2 := db.InsertMuscle(m)

	if err2 != nil {
		return 400, "Error to attemp create Muscle"
	}

	return int(result), "Muscle created"
}

func UpdateMuscle(body string, user string, id int) (int, string) {
	var m models.Muscle
	
	err := json.Unmarshal([]byte(body), &m)

	if err != nil {
		return 400, "Error to attemp convert to JSON "+err.Error()
	}

	isAdmin, msg := db.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	m.Id_Muscle = id
	err2 := db.UpdateMuscle(m)

	if err2 != nil {
		return 400, "Error to update muscle with ID: "+strconv.Itoa(m.Id_Muscle)
	}

	return 200, "Muscle updated"
}

func DeleteMuscle(user string, id int) (int, string) {
	isAdmin, msg := db.IsAdmin(user)

	if !isAdmin {
		return 400, msg
	}

	err2 := db.DeleteMuscle(id)

	if err2 != nil {
		return 400, "Error to Delete Muscle"
	}

	return 200, "Muscle deleted"
}

func SelectMuscle(request events.APIGatewayV2HTTPRequest) (int, string) {
	var m models.Muscle
	var orderType, orderField string

	param := request.QueryStringParameters

	orderType = param["orderType"]
	orderField = param["orderField"]

	if !strings.Contains("ITDFPCS", orderField) {
		orderField=""
	}

	var choice string
	
}