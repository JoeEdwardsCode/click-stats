package clickstatsservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type ClickEvent struct {
	Shape    string `json:"shape"`
	Color    string `json:"color"`
	Quadrant string `json:"quadrant"`
}

func RecordClickEvent(ctx context.Context, event events.APIGatewayProxyRequest) ([]byte, error) {
	var applicationError error
	var clickEvent ClickEvent
	eventBody := []byte(event.Body)

	applicationError = json.Unmarshal(eventBody, &clickEvent)

	if applicationError != nil {
		return nil, applicationError
	}

	// Placeholder for dynamo PUT
	log.Print(fmt.Sprintf("Shape: %s, Color: %s Quadrant: %s", clickEvent.Shape, clickEvent.Color, clickEvent.Quadrant))

	body, applicationError := json.Marshal(map[string]interface{}{
		"message": "click recorded",
	})

	return body, nil
}
