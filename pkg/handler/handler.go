package handler

import (
	"fmt"
	"terraform-lambda-dynamodb/pkg/domain/model"
)

func Handler(event model.MyEvent) (model.MyResponse, error) {
	return model.MyResponse{Message: fmt.Sprintf("Hello %s!!", event.Key1)}, nil
}
