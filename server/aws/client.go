package aws

import (
	"fmt"
	"os"

	"goscrum/goscrum/server/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/jinzhu/gorm"
)

var dbClient *gorm.DB

func NewDBClient(debug bool) *gorm.DB {
	if dbClient != nil {
		return dbClient
	}

	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True`,
		os.Getenv(constants.DBUsername),
		os.Getenv(constants.DBPassword),
		os.Getenv(constants.DatabaseHostName),
		os.Getenv(constants.DatabasePort),
		os.Getenv(constants.DatabaseName))

	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(debug)
	dbClient = db
	return dbClient
}

func NewSession() client.ConfigProvider {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(constants.AwsRegionName))},
	)
	if err != nil {
		panic(err)
	}
	return sess
}

func NewCognitoClient(sess client.ConfigProvider) *cognitoidentityprovider.CognitoIdentityProvider {
	return cognitoidentityprovider.New(sess)
}
