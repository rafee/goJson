package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Person object holds the name and id of a random person
type Person struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := rand.Int()
	name := "Stupid Rafee"
	p := Person{Name: name, ID: id}
	b, err := json.Marshal(p)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}
	return resp, err
}

func main() {
	lambda.Start(handler)
}
