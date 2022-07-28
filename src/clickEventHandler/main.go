package main

import (
	"context"

	"click-stats/src/clickStatsService"
	"click-stats/src/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ClickEventHandler(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	applicationError := clickStatsService.RecordClickEvent(event)

	if applicationError != nil {
		errorBuffer := []byte(applicationError.Error())
		response = utils.ClickStatsResponse(500, errorBuffer)
		return response, nil
	}

	responseText := "click recorded"
	response = utils.ClickStatsResponse(200, []byte(responseText))
	return response, nil
}

func main() {
	lambda.Start(ClickEventHandler)
}
