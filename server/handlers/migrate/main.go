package main

import (
	"context"
	"goscrum/goscrum/server/db"

	"goscrum/goscrum/server/aws"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func HandleRequest(_ context.Context) error {
	dbClient := aws.NewDBClient(true)
	defer dbClient.Close()

	err := dbClient.AutoMigrate(&db.User{}, db.Organization{}, db.OrganizationUser{}).Error

	if err != nil {
		panic(err)
	}

	return err
}

func main() {
	lambda.Start(HandleRequest)
}
