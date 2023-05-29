package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"terraform-lambda-dynamodb/pkg/config"
	"terraform-lambda-dynamodb/pkg/domain/entity"
	"terraform-lambda-dynamodb/pkg/io"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
)

type MyEvent struct {
	Key1 string `json:"key1"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func hello(event MyEvent) (MyResponse, error) {

	return MyResponse{Message: fmt.Sprintf("Hello %s!!", event.Key1)}, nil
}

func main() {
	ctx := context.Background()

	if os.Getenv("ENV") != "dev" {
		lambda.Start(hello)

	} else {
		cfg, err := config.LoadConfig(ctx)
		if err != nil {
			log.Fatal(err)
		}
		db, err := io.NewDb(cfg)
		if err != nil {
			log.Fatal(err)
		}

		// テーブル作成をする為に、一度テーブルを削除します
		db.Table("UserTable").DeleteTable().Run()

		err = db.CreateTable("UserTable", entity.User{}).Run()
		if err != nil {
			panic(err)
		}
		// テーブルの指定
		table := db.Table("UserTable")

		// User構造体をuser変数に定義
		var user entity.User

		// DBにPutします
		err = table.Put(&entity.User{UserID: "1234", Name: "太郎", Age: 20, Text: "text1"}).Run()
		if err != nil {
			panic(err)
		}
		err = table.Get("UserID", "1234").Range("Name", dynamo.Equal, "太郎").One(&user)
		if err != nil {
			panic(err)
		}
		fmt.Printf("GetDB%+v\n", user)

		// DBのデータをUpdateします
		text := "新しいtextです"
		err = table.Update("UserID", "1234").Range("Name", "太郎").Set("Text", text).Value(&user)
		if err != nil {
			panic(err)
		}
		fmt.Printf("UpdateDB%+v\n", user)

		// DBのデータをDeleteします
		err = table.Delete("UserID", "1").Range("Name", "Test1").Run()
		if err != nil {
			panic(err)
		}

		// Delete出来ているか確認
		err = table.Get("UserID", "1").Range("Name", dynamo.Equal, "Test1").One(&user)
		if err != nil {
			// Delete出来ていれば、dynamo: no item found のエラーとなる
			fmt.Println("getError:", err)
		}

	}
}
