package command

import (
	"context"
	"fmt"
	"log"
	"os"
	"terraform-lambda-dynamodb/pkg/config"
	"terraform-lambda-dynamodb/pkg/domain/entity"
	"terraform-lambda-dynamodb/pkg/handler"
	"terraform-lambda-dynamodb/pkg/io"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
)

const dynamoDbTableName = "example_lambda_dynamodb_table"

func Run() {
	run(context.Background())
}

func run(ctx context.Context) {
	// prod
	if os.Getenv("ENV") != "dev" {
		lambda.Start(handler.Handler)

	} else {
		// dev
		cfg, err := config.LoadConfig(ctx)
		if err != nil {
			log.Fatal(err)
		}
		db, err := io.NewDb(cfg)
		if err != nil {
			log.Fatal(err)
		}

		// テーブル作成をする為に、一度テーブルを削除します
		db.Table(dynamoDbTableName).DeleteTable().Run()

		err = db.CreateTable(dynamoDbTableName, entity.User{}).Run()
		if err != nil {
			panic(err)
		}
		// テーブルの指定
		table := db.Table(dynamoDbTableName)

		// User構造体をuser変数に定義
		var user entity.User

		// DBにPutします
		err = table.Put(&entity.User{ID: "1234", Name: "太郎"}).Run()
		if err != nil {
			panic(err)
		}
		err = table.Get("Id", "1234").Range("Name", dynamo.Equal, "太郎").One(&user)
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetDB%+v\n", user)

		// DBのデータをUpdateします
		text := "新しいtextです"
		err = table.Update("Id", "1234").Range("Name", "太郎").Set("Text", text).Value(&user)
		if err != nil {
			panic(err)
		}
		fmt.Printf("UpdateDB%+v\n", user)

		// DBのデータをDeleteします
		err = table.Delete("Id", "1234").Range("Name", "Test1").Run()
		if err != nil {
			panic(err)
		}

		// Delete出来ているか確認
		err = table.Get("Id", "1234").Range("Name", dynamo.Equal, "新しいtextです").One(&user)
		if err != nil {
			// Delete出来ていれば、dynamo: no item found のエラーとなる
			fmt.Println("getError:", err)
		}

	}
}
