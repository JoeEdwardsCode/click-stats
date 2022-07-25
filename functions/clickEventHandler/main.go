package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/JoeEdwardsCode/click-stats/utils"
)

func ClickEventHandler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	var applicationError error

	body, applicationError := json.Marshal(map[string]interface{}{
		"message": "click recorded",
	})

	if applicationError != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, applicationError
	}

	response := utils.ClickStatsResponse(200, body)

	return response, nil
}

func main() {
	lambda.Start(ClickEventHandler)
}
