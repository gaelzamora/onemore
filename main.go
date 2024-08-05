package main

import (
	"context"
	"log"
	"onemore/awsgo"
	"onemore/db"
	"onemore/handlers"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// The main function, execute all Lambda and logical in the API
func main() {
	lambda.Start(ExecuteLambda)
}

// Execute Lambda
func ExecuteLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	log.Println("Start Lambda")

	// Call StartAWS to configure AWS region and other stuffs
	awsgo.StartAWS()

	// Call function
	if !IsParamValid() {
		panic("Error in parameters, you has send param")
	}
	
	// 
	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	db.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headersResp := map[string]string {
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body: string(message),
		Headers: headersResp,
	}

	return res, nil
}

// Review params that are sent by Lambda
func IsParamValid() bool {
	_, getParam := os.LookupEnv("SecretName")
	if !getParam {
		return getParam
	}
	
	_, getParam = os.LookupEnv("UrlPrefix")
	if !getParam {
		return getParam
	}

	return getParam
}