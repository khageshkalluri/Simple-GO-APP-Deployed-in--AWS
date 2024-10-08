package main

import (
	"dynamic/app"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Details struct {
	Username string `json:"username"`
}

func Operation(details Details) (string, error) {
	if details.Username == "" {
		return "", fmt.Errorf("the username should not be null")
	}
	return fmt.Sprintf("Welcome to the AWS %s", details.Username), nil
}

func main() {
	newapp := app.NewApp()
	lambda.Start(newapp.Api.RegisterClient)
}
