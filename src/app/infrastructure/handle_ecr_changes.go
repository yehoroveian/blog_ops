package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"gitlab.com/blog/ops/pkg/log"
)

func (s *Service) HandleDeployECRChanges(ctx context.Context, event events.EventBridgeEvent) error {
	detail := struct {
		ActionType     string `json:"action-type"`
		Result         string `json:"result"`
		RepositoryName string `json:"repository-name"`
		ImageTag       string `json:"image-tag"`
	}{}

	if event.Detail == nil {
		log.Warnf("handleTestRequest details are nil")

		return nil
	}

	if err := json.Unmarshal(event.Detail, &detail); err != nil {
		return fmt.Errorf("failed to unmarshal detail: %w", err)
	}

	imageURI := fmt.Sprintf("%s/%s:%s", s.imageURI, detail.RepositoryName, s.imageTag)

	lambdaName, err := s.db.FetchECRLambdaNamespace(ctx, detail.RepositoryName)
	if err != nil {
		return fmt.Errorf("failed to fetch lambda namespace: %w", err)
	}

	if lambdaName != "" {
		_, err = s.cl.UpdateFunctionCode(ctx, &lambda.UpdateFunctionCodeInput{
			FunctionName: &lambdaName,
			ImageUri:     &imageURI,
			Publish:      true,
		})
		if err != nil {
			return err
		}

		return s.db.UpdateECRLambdaNamespace(ctx, detail.RepositoryName)
	}

	lambdaName = convertECRImageNameToLambdaName(detail.RepositoryName)

	_, err = s.cl.CreateFunction(ctx, &lambda.CreateFunctionInput{
		PackageType:   types.PackageTypeImage,
		Code:          &types.FunctionCode{ImageUri: &imageURI},
		FunctionName:  &lambdaName,
		Role:          &s.lambdaRunnerRole,
		Architectures: []types.Architecture{types.ArchitectureArm64},
		Publish:       true,
	})
	if err != nil {
		return fmt.Errorf("failed to create lambda function: %w", err)
	}

	err = s.db.SaveECRLambdaNamespace(ctx, detail.RepositoryName, lambdaName)
	if err != nil {
		return fmt.Errorf("failed to save lambda namespace: %w", err)
	}

	return nil
}

func convertECRImageNameToLambdaName(imageName string) string {
	var lambdaName string

	lowCaseParts := strings.Split(imageName, "-")
	for i := range lowCaseParts {
		lambdaName += strings.ToUpper(string(lowCaseParts[i][0])) + lowCaseParts[i][1:]
	}

	return lambdaName
}
