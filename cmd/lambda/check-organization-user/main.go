package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/valyala/fasthttp"
	"grafana-api/cmd/lambda/check-organization-user/handler"
)

var lambdaHandler handler.LambdaHandler

func init() {
}

func main() {
	lambda.Start(lambdaHandler.HandleRequest)
}
