package ecr

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (d *DynamoRepository) FetchECRLambdaNamespace(ctx context.Context, containerName string) (string, error) {
	resp, err := d.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key:             map[string]types.AttributeValue{"ecr_container_name": &types.AttributeValueMemberS{Value: containerName}},
		TableName:       &tableName,
		AttributesToGet: []string{"lambda_name"},
	})
	if err != nil {
		return "", fmt.Errorf("failed to fetch ecrl lambda namespace: %w", err)
	}

	if resp == nil || len(resp.Item) == 0 {
		return "", nil
	}

	if val, ok := resp.Item["lambda_name"]; ok {
		if lambdaName, ok := val.(*types.AttributeValueMemberS); ok {
			return lambdaName.Value, nil
		}
	}

	return "", nil
}
