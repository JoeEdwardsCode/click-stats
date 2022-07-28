package clickStatsService

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ClickEvent struct {
	Shape    string `json:"shape"`
	Color    string `json:"color"`
	Quadrant string `json:"quadrant"`
}

type ClickRecord struct {
}

func getClickEvent(eventBody []byte) (ClickEvent, error) {
	var clickEvent ClickEvent
	json.Unmarshal(eventBody, &clickEvent)

	if clickEvent.Shape == "" || clickEvent.Color == "" || clickEvent.Quadrant == "" {
		return clickEvent, errors.New("request body invalid")
	}

	return clickEvent, nil
}

func recordClick(clickEvent ClickEvent) ([]byte, error) {
	input := &dynamodb.PutItemInput{
		Item:      nil,
		TableName: "clickEvents",
	}

	// config, configError := config.LoadDefaultConfig(context.TODO())
	// if configError != nil {
	// 	return nil, configError
	// }

	// client := dynamodb.NewFromConfig(config)

	// response, putError := client.PutItem(context.TODO(), input)
	return nil, nil
}

func RecordClickEvent(context context.Context, event events.APIGatewayProxyRequest) ([]byte, error) {
	var applicationError error
	eventBody := []byte(event.Body)

	clickEvent, applicationError := getClickEvent(eventBody)

	if applicationError != nil {
		return nil, applicationError
	}

	recordClickEventResponse, applicationError := recordClick(clickEvent)

	return recordClickEventResponse, nil
}
