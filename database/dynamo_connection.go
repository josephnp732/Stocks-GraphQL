package database

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var db *dynamo.DB
var awsRegion = "us-east-1"
var tableName = "stocks"
var awsAccessToken = "Add_Your_Token_Here"
var awsSecret = "Add_Your_Secret_Here"

// Table gets the dynamo table after connection
var Table dynamo.Table

func init() {

	if aT, ok := os.LookupEnv(("AWS_ACCESS_TOKEN")); ok {
		awsAccessToken = aT
	}

	if s, ok := os.LookupEnv(("AWS_SECRET")); ok {
		awsSecret = s
	}

	// Establish a new session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		log.Panicln(err)
	}

	// Add AWS Credentials
	awsCredentials := credentials.NewStaticCredentials(awsAccessToken, awsSecret, "")

	// New DynamoDB connection
	db = dynamo.New(sess, &aws.Config{Region: aws.String(awsRegion), Credentials: awsCredentials})
	Table = db.Table(tableName)
}
