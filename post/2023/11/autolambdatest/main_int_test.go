//go:build integration_test

package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	awsSSM "github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	pulumiAws "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
	"time"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AutoLambdaIntTestSuite struct {
	suite.Suite
	stack            *auto.Stack
	subjectParamPath string
	roleArn          string
}

func (suite *AutoLambdaIntTestSuite) SetupSuite() {

	fmt.Println("BeforeSuite")

	ctx := context.TODO()
	projectName := "autolambdatest"
	stackName := "PoC"

	stack, err := auto.UpsertStackInlineSource(ctx, stackName, projectName, deployFunc)
	if err != nil {
		log.Panicf("Unable to get pulumi stack: %s", err)
	}

	// wire up our update to stream progress to stdout
	stdoutStreamer := optup.ProgressStreams(os.Stdout)

	res, err := stack.Up(ctx, stdoutStreamer)
	if err != nil {
		log.Panicf("Unable to up pulumi stack: %s", err)
	}

	fmt.Printf("outputs: %+v", res.Outputs)
	suite.stack = &stack
	suite.subjectParamPath = res.Outputs["subjectParamPath"].Value.(string)
	suite.roleArn = res.Outputs["roleArn"].Value.(string)
	time.Sleep(1 * time.Second) // Wait for AWS to catch up
}

func (suite *AutoLambdaIntTestSuite) TearDownSuite() {
	fmt.Println("AfterSuite")
	_, err := suite.stack.Destroy(context.TODO())
	if err != nil {
		log.Panicf("Unable to destroy pulumi stack: %s", err)
	}
}

func deployFunc(ctx *pulumi.Context) error {

	current, err := pulumiAws.GetCallerIdentity(ctx, nil, nil)
	if err != nil {
		return err
	}
	assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect: pulumi.StringRef("Allow"),
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type: "Service",
						Identifiers: []string{
							"lambda.amazonaws.com",
						},
					},
					{ // Allow any "user" with permissions to assume this role
						Type: "AWS",
						Identifiers: []string{
							fmt.Sprintf("arn:aws:iam::%s:root", current.AccountId),
						},
					},
				},
				Actions: []string{
					"sts:AssumeRole",
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	roleForLambda, err := iam.NewRole(ctx, "autolambdatest", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(assumeRole.Json),
	})

	if err != nil {
		return err
	}

	subjectParam, err := ssm.NewParameter(ctx, "/autolambdatest/subject", &ssm.ParameterArgs{
		Type:  pulumi.String("String"),
		Value: pulumi.String("World"),
	})
	if err != nil {
		return err
	}

	subjectParamPerms := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
		Statements: iam.GetPolicyDocumentStatementArray{iam.GetPolicyDocumentStatementArgs{
			Effect: pulumi.String("Allow"),
			Actions: pulumi.StringArray{
				pulumi.String("ssm:GetParameter"),
			},
			Resources: pulumi.StringArray{
				subjectParam.Arn,
			},
		}},
	})

	_, err = iam.NewRolePolicy(ctx, "autolambdatestSubjectParam", &iam.RolePolicyArgs{
		Role:   roleForLambda.ID(),
		Policy: subjectParamPerms.Json(),
	})

	ctx.Export("subjectParamPath", subjectParam.Name)
	ctx.Export("roleArn", roleForLambda.Arn)

	return nil
}

func (suite *AutoLambdaIntTestSuite) getLambdaRoleCfg() aws.Config {
	log.Println("Setup")

	roleToAssume := suite.roleArn

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		log.Fatal("error: ", err)
	}
	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN using the MFA token code provided.
	creds := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(cfg), roleToAssume)

	cfg.Credentials = aws.NewCredentialsCache(creds)
	return cfg
}

func (suite *AutoLambdaIntTestSuite) TestHandleRequest() {

	cfg := suite.getLambdaRoleCfg()
	ssmClient := awsSSM.NewFromConfig(cfg)

	output := handleRequest(ssmClient, suite.subjectParamPath)

	expectedOutput := "Hello World"
	assert.Equal(suite.T(), expectedOutput, output)
}

func TestAutoLambdaIntTestSuite(t *testing.T) {
	suite.Run(t, new(AutoLambdaIntTestSuite))
}
