package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Person struct {
	Name string `json:"name"`
	id   int    `json:"id"`
}

func getJsonHandler(ctx context.Context, r events.APIGatewayProxyRequest) (error, events.APIGatewayProxyResponse) {
	id := rand.Int()
	name := "Stupid Rafee"
	p := Person{Name: name, id: id}
	b, err := json.Marshal(p)
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}
	return err, resp
}

func main() {
	lambda.Start(getJsonHandler)
}
