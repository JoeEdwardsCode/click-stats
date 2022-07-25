package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/JoeEdwardsCode/click-stats/clickStatsService"
	"github.com/JoeEdwardsCode/click-stats/utils"
)

func ClickEventHandler(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var applicationError error
	var response events.APIGatewayProxyResponse

	serviceResponse, applicationError := clickStatsService.RecordClickEvent(context, event)

	if applicationError != nil {
		errorBuffer := []byte(applicationError.Error())
		response = utils.ClickStatsResponse(400, errorBuffer)
		return response, nil
	}

	response = utils.ClickStatsResponse(200, serviceResponse)
	return response, nil
}

func main() {
	lambda.Start(ClickEventHandler)
}
