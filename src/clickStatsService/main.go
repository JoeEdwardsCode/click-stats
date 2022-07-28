package clickStatsService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ClickEvent struct {
	Shape    string `json:"shape"`
	Color    string `json:"color"`
	Quadrant string `json:"quadrant"`
}

type ClickEventTableRecord struct {
	record_type string
	timestamp   string
	shape       string
	color       string
	quadrant    string
}

func getClickEvent(eventBody []byte) (ClickEvent, error) {
	var clickEvent ClickEvent
	json.Unmarshal(eventBody, &clickEvent)

	if clickEvent.Shape == "" || clickEvent.Color == "" || clickEvent.Quadrant == "" {
		return clickEvent, errors.New("request body invalid")
	}

	return clickEvent, nil
}

func RecordClickEvent(event events.APIGatewayProxyRequest) error {
	clickEvent, eventError := getClickEvent([]byte(event.Body))

	if eventError != nil {
		return eventError
	}

	var tableName string = os.Getenv("TABLE_NAME")
	tableRecord := ClickEventTableRecord{
		record_type: "CLICK",
		timestamp:   time.Now().UTC().String(),
		color:       clickEvent.Color,
		shape:       clickEvent.Shape,
		quadrant:    clickEvent.Quadrant,
	}

	itemMap := map[string]types.AttributeValue{
		"record_type": &types.AttributeValueMemberS{Value: tableRecord.record_type},
		"timestamp":   &types.AttributeValueMemberS{Value: tableRecord.timestamp},
		"shape":       &types.AttributeValueMemberS{Value: tableRecord.shape},
		"quadrant":    &types.AttributeValueMemberS{Value: tableRecord.quadrant},
		"color":       &types.AttributeValueMemberS{Value: tableRecord.color},
	}

	itemInput := &dynamodb.PutItemInput{
		Item:      itemMap,
		TableName: &tableName,
	}

	config, configError := config.LoadDefaultConfig(context.TODO())
	if configError != nil {
		return configError
	}

	fmt.Printf("Writing item:\n%+v", tableRecord)
	client := dynamodb.NewFromConfig(config)
	_, putError := client.PutItem(context.TODO(), itemInput)

	if putError != nil {
		return putError
	}

	return nil
}
