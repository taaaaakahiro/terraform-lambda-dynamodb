package handler

import (
	"context"
	"fmt"
	"log"
	"terraform-lambda-dynamodb/pkg/config"
	"terraform-lambda-dynamodb/pkg/domain/entity"
	"terraform-lambda-dynamodb/pkg/domain/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func Handler(event model.MyEvent) (model.MyResponse, error) {

	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String(cfg.DB.AwsRegion)})
	table := db.Table(cfg.DB.DynamoTableName)
	var user entity.User
	err = table.Get("UserId", "001").Range("Name", dynamo.Equal, "テストデータ1").One(&user)
	if err != nil {
		log.Fatal(err)
	}

	return model.MyResponse{Message: fmt.Sprintf("Hello %s:%s!!", user.Name, user.Text)}, nil
}
