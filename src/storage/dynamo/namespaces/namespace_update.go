package namespaces

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (d *DynamoRepository) UpdateECRLambdaNamespace(ctx context.Context, containerName string) error {
	updateExpression := "SET action_type = :action_type, updated_at = :updated_at"

	input := &dynamodb.UpdateItemInput{
		TableName:        &tableName,
		Key:              map[string]types.AttributeValue{"ecr_container_name": &types.AttributeValueMemberS{Value: containerName}},
		UpdateExpression: &updateExpression,
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":action_type": &types.AttributeValueMemberS{Value: actionTypeUpdate},
			":updated_at":  &types.AttributeValueMemberS{Value: time.Now().Format("2006-01-02T15:04:05Z")},
		},
	}

	_, err := d.db.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update lambda namespace: %w", err)
	}

	return nil
}
