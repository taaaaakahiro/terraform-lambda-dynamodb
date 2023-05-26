package io

import (
	"terraform-lambda-dynamodb/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewDb(cfg *config.Config) (*dynamo.DB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.DB.AwsRegion),
		Endpoint:    aws.String(cfg.DB.DynamoEndpoint),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		return nil, err
	}
	db := dynamo.New(sess)

	return db, nil
}
