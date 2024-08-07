package infrastructure

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	l "github.com/aws/aws-sdk-go-v2/service/lambda"

	"gitlab.com/blog/ops/src/storage/dynamo/namespaces"
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

	cl *l.Client
	db namespaces.Repository
}

// Interface compliance check.
var _ Provider = (*Service)(nil)

func New(imageURI, imageTag, secretName, secretRegion, lambdaRunnerRole string, awsClient *l.Client, db namespaces.Repository) *Service {
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
