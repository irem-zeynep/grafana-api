package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	_ "github.com/valyala/fasthttp"
	"grafana-api/cmd/lambda/check-organization-user/handler"
	"grafana-api/domain"
	"grafana-api/infrastructure/event/sns"
	"grafana-api/infrastructure/generator/password"
	"grafana-api/infrastructure/http/grafana"
	"grafana-api/infrastructure/persistence/timestream"
	"grafana-api/infrastructure/secretmanager"
)

var lambdaHandler handler.LambdaHandler

func init() {
	lambdaHandler.Logger = logrus.New()
	lambdaHandler.Logger.SetFormatter(&logrus.JSONFormatter{})

	secret := secretmanager.Init()

	client := grafana.NewClient(secret.GrafanaClient)

	publisher := sns.NewEventPublisher(secret.ErrorTopic)

	auditRepo := timestream.NewAuditRepository(secret.TimeStreamDB)

	passwordGenerator := password.NewPasswordGenerator()

	lambdaHandler.Serv = domain.NewGrafanaService(client, passwordGenerator, lambdaHandler.Logger)
	lambdaHandler.ExceptionServ = domain.NewExceptionService(publisher, lambdaHandler.Logger)
	lambdaHandler.AuditServ = domain.NewAuditService(auditRepo, lambdaHandler.Logger)
}

func main() {
	lambda.Start(lambdaHandler.HandleRequest)
}
