package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	l "github.com/aws/aws-sdk-go-v2/service/lambda"

	"gitlab.com/blog/ops/src/app/infrastructure"
	"gitlab.com/blog/ops/src/config"
	"gitlab.com/blog/ops/src/storage/dynamo/namespaces"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.NewDeploy()
	if err != nil {
		return fmt.Errorf("load api config: %w", err)
	}

	awsConfig, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	awsClient := l.NewFromConfig(awsConfig)
	dynamoClient := namespaces.New(dynamodb.NewFromConfig(awsConfig))

	svc := infrastructure.New(cfg.ECR.URI, cfg.ECR.User, cfg.Secret.SecretName, cfg.Secret.SecretRegion, cfg.Lambda.LambdaRunnerRole, awsClient, dynamoClient)

	lambda.StartWithOptions(svc.HandleDeployECRChanges, lambda.WithContext(ctx))

	return nil
}
