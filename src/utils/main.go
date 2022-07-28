package utils

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ClickStatsResponse(statusCode int, body []byte) events.APIGatewayProxyResponse {
	var key = "data"

	if statusCode != 200 {
		key = "error"
	}

	jsonResponse, _ := json.Marshal(map[string]string{key: string(body)})
	var buffer bytes.Buffer
	json.HTMLEscape(&buffer, jsonResponse)

	return events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            buffer.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
