package handlers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status 

	stringBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("Error while marshaling response body: %v", err)

		errResp := events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: `"Message": "Internal server error"`,
			Headers: map[string]string{"Content-Type": "application/json"},
		}

		return &errResp, err
	}

	resp.Body = string(stringBody)

	return &resp, nil
}




