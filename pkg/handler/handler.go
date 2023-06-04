package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"terraform-lambda-dynamodb/pkg/config"
	"terraform-lambda-dynamodb/pkg/domain/entity"
	"terraform-lambda-dynamodb/pkg/domain/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func Handler(event model.MyEvent) (events.APIGatewayProxyResponse, error) {
	log.Println("lambda start")
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Println("failed to load config")
		log.Fatal(err)
	}
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String(cfg.DB.AwsRegion)})
	table := db.Table(cfg.DB.DynamoTableName)
	log.Println("succes to select table")

	var user entity.User
	intUserID, _ := strconv.Atoi(event.Key1)
	id := fmt.Sprintf("%03d", intUserID)
	err = table.Get("UserId", id).Range("Name", dynamo.Equal, "テストデータ"+event.Key1).One(&user)
	if err != nil {
		log.Println("failed to get config")
		log.Fatal(err)
	}

	log.Println("end lambda")

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello %s:%s!!", user.Name, user.Text),
		StatusCode: 200,
	}, nil
	// return model.MyResponse{Message: fmt.Sprintf("Hello %s:%s!!", user.Name, user.Text)}, nil
}
