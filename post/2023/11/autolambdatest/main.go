package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMGetParameterAPI interface {
	GetParameter(ctx context.Context,
		params *ssm.GetParameterInput,
		optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

func handleRequest(api SSMGetParameterAPI, subjectNameParamPath string) string {

	return ""

}
