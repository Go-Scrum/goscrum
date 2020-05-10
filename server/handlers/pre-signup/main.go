package main

import (
	"goscrum/goscrum/server/constants"
	"goscrum/goscrum/server/db"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

func handler(event *events.CognitoEventUserPoolsPreSignup) (*events.CognitoEventUserPoolsPreSignup, error) {
	if event.UserPoolID == os.Getenv(constants.UserPoolID) {
		user := db.User{
			FirstName: event.Request.UserAttributes["given_name"],
			LastName:  event.Request.UserAttributes["family_name"],
			Email:     event.Request.UserAttributes["email"],
		}
		spew.Dump(user)
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
