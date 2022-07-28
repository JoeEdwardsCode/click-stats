package utils

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ClickStatsResponse(statusCode int, body []byte) events.APIGatewayProxyResponse {

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	return events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
