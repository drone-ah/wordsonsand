package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMGetParameterAPI interface {
	GetParameter(ctx context.Context,
		params *ssm.GetParameterInput,
		optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

func handleRequest(api SSMGetParameterAPI, subjectNameParamPath string) string {

	subjectNameParam, err := api.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: &subjectNameParamPath,
	})

	if err != nil {
		return fmt.Sprintf("Error Accessing Parameter: %s", err)
	}

	return fmt.Sprintf("Hello %s", *subjectNameParam.Parameter.Value)

}
