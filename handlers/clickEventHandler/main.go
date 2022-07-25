package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/JoeEdwardsCode/click-stats/utils"
)

func ClickEventHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var applicationError error
	var response events.APIGatewayProxyResponse

	body, applicationError := json.Marshal(map[string]interface{}{
		"message": "click recorded",
	})

	if applicationError != nil {
		errorBuffer := []byte(applicationError.Error())
		response = utils.ClickStatsResponse(400, errorBuffer)
		return response, nil
	}

	response = utils.ClickStatsResponse(200, body)
	return response, nil
}

func main() {
	lambda.Start(ClickEventHandler)
}
