package namespaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=ecr Repository

var tableName = "ECRLambdasNamespaces"

const (
	actionTypeUpdate = "update"
	actionTypeCreate = "create"
)

type Repository interface {
	SaveECRLambdaNamespace(ctx context.Context, containerName, lambdaName string) error
	FetchECRLambdaNamespace(ctx context.Context, containerName string) (string, error)
	UpdateECRLambdaNamespace(ctx context.Context, containerName string) error
}

// Interface compliance check.
var _ Repository = (*DynamoRepository)(nil)

type DynamoRepository struct {
	db *dynamodb.Client
}

func New(db *dynamodb.Client) *DynamoRepository {
	return &DynamoRepository{
		db: db,
	}
}
