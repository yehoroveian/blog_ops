package infrastructure

import (
	"context"

	"gitlab.com/blog/ops/src/storage/dynamo/ecr"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type Provider interface {
	HandleDeployECRChanges(ctx context.Context, event events.EventBridgeEvent) error
}

type Service struct {
	imageURI         string
	imageTag         string
	secretName       string
	secretRegion     string
	lambdaRunnerRole string

	cl *lambda.Client
	db ecr.Repository
}

// Interface compliance check.
var _ Provider = (*Service)(nil)

func New(imageURI, imageTag, secretName, secretRegion, lambdaRunnerRole string, awsClient *lambda.Client, db ecr.Repository) *Service {
	return &Service{
		imageURI:         imageURI,
		imageTag:         imageTag,
		secretName:       secretName,
		secretRegion:     secretRegion,
		lambdaRunnerRole: lambdaRunnerRole,

		cl: awsClient,
		db: db,
	}
}
