package namespaces

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (d *DynamoRepository) SaveECRLambdaNamespace(ctx context.Context, containerName, lambdaName string) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"ecr_container_name": &types.AttributeValueMemberS{Value: containerName},
			"lambda_name":        &types.AttributeValueMemberS{Value: lambdaName},
			"action_type":        &types.AttributeValueMemberS{Value: actionTypeCreate},
			"updated_at":         &types.AttributeValueMemberS{Value: time.Now().Format("2006-01-02T15:04:05Z")},
		},
	}

	_, err := d.db.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to save ECRL lambda namespace: %w", err)
	}

	return nil
}
