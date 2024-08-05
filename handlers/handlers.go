package handlers

import (
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/internal/auth"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOK, statusCode, user := validAuthorization(path, method, headers)

	if !isOK {
		return statusCode, user
	}

	if len(path) < 4 {
		return 400, "Path too short"
	}

	switch path[1:5] {
	case "user":
		return UserToProcess(body, path, method, user, id, request)
	case "musc":
		return MuscleToProcess(body, path, method, user, idn, request)
	case "rout":
		return RoutineToProcess(body, path, method, idn, request)
	case "exer":
		return ExerciseToProcess(body, path, method, idn, request)
	case "seri":
		return SerieToProcess(body, path, method, idn, request)
	}
}

func validAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	token := headers["authorization"]

	if len(token)==0 {
		return false, 401, "Required Token"
	}

	allOK, err, msg := auth.ValidToken(token)

	if !allOK {
		if err != nil {
			return false, 401, err.Error()
		} else {
			log.Println("Message error: "+msg)
			return false, 401, msg
		}
	}

	return true, 200, msg
}

func UserToProcess(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	if path == "/users/me" {
		switch method {
		case "GET":
			return routers.GetUser(body, user)
		case "PUT":
			return routers.UpdateUser(body, user)
		}
	}

	if path == "/users" {
		if method == "GET" {
			switch method {
			case "GET":
				return routers.SelectUsers(body, user, request)
			}
		}
	}
}

func MuscleToProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertMuscle(body, user)
	case "PUT":
		return routers.UpdateMuscle(body, user, id)
	case "DELETE":
		return routers.DeleteMuscle(user, id)
	case "GET":
		return routers.SelectMuscle(request)
	}

	return 400, "Method invalid"
}