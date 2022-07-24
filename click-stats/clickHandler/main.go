package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func ClickHandler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, applicationError := json.Marshal(map[string]interface{}{
		"message": "click recorded",
	})
	if applicationError != nil {
		return Response{StatusCode: 400}, applicationError
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(ClickHandler)
}
