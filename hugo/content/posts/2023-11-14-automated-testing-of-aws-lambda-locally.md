---
categories: []
date: "2023-11-14T14:23:37Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:06"
  _last_editor_used_jetpack: block-editor
  _publicize_job_id: "89358302619"
  timeline_notification: "1699971818"
  wordads_ufa: s:wpcom-ufa-v4:1700043369
parent_id: "0"
password: ""
status: publish
tags:
- golang
- integration testing
- pulumi
- testing
title: How to run your lambda code locally as its role (for testing)
url: /2023/11/14/automated-testing-of-aws-lambda-locally/
---

While AWS Lambda is fantastic in providing a serverless platform with few
worries about maintaining servers, it is not the easiest to test in an automated
fashion with rapid feedback.

You could write end to end tests, but it means a deployment after each change
and then checking the logs to see what failed. Even if you use iac
(terraform/pulumi), the deployment will take seconds or a minute or two - not
exact rapid test feedback.

What I have been doing is to set up a hook which is called from the lambda
handler, which can also be called locally. Within the test, I then assume the
role that runs the lambda and then test the hook.

This mechanism allows to me easily test that the permissions are set up
correctly and that details are in place for the code to work.

For the full end to end test, I then have a simple smoke test or two.

The code samples are in golang(only because it happens to be my current language
of choice), but the idea should be equally applicable in other languages.

<!--more-->

# Assuming The Role {#assuming-the-role}

```go
  roleToAssume := os.Getenv("AUTH_LAMBDA_ROLE_ARN")

    ctx := context.TODO()
    cfg, err := config.LoadDefaultConfig(ctx)

    if err != nil {
        logger.Fatal("error: ", err)
    }
    // Create the credentials from AssumeRoleProvider to assume the role
    // referenced by the "myRoleARN" ARN using the MFA token code provided.
    creds := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(cfg), roleToAssume)

    logger.Debugf("creds: %v", creds)

    cfg.Credentials = aws.NewCredentialsCache(creds)
```

the cfg is then passed into the `New` method for the resource you are interested
in. e.g.:

```go
   ssmClient := ssm.NewFromConfig(cfg)
```

# Working Example {#working-example}

You can find a full, working example test in
[my github repo](https://github.com/drone-ah/wordsonsand) under
[post/2023/11/autolambdatest](https://github.com/drone-ah/wordsonsand/tree/main/post/2023/11/autolambdatest)

NOTE: It WILL automatically try and deploy a role and a ssm parameter, and it
will delete it after the test.

The `BeforeSuite` will deploy the minimum configuration to be able to run the
test, and the `AfterSuite` will destroy the same stack.

You will likely need to log into pulumi to get this test to work.

If you run into permissions issue for AssumeRole, read on.

# AssumeRole Permissions {#assumerole-permissions}

For this to work, the user running the tests need to have permissions to
`AssumeRole`.

There are two steps to this. The first part is to allow \"anyone\" to
`AssumeRole` the relevant role:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::<your-account-id>:root"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

This will allow \"any user\" to assume the role, as long as they have the
permission to do so.

`arn:aws:iam::[your-account]:root` is a special user the represents the account
(and the non-IAM root user). Since IAM (user, roles etc.) exists under the
\"root\" account, all calls are also authenticated by this account - i.e. all
users, roles etc. in IAM is also this account. There is
[a post on reddit discussing what exactly the root iam principal is for more information](https://www.reddit.com/r/aws/comments/oorjl2/what_exactly_is_the_root_iam_principal/)

Finally, unless you have the `Administrator`Access policy set against your
account, you will also need to attach a policy to the relevant group (or your
user) that grants permissions to call `sts:AssumeRole` (or `*`)

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "123",
      "Effect": "Allow",
      "Action": ["sts:AssumeRole"],
      "Resource": ["arn:aws:iam::123456789012:role/desired-role"]
    }
  ]
}
```

You can of course, also use `*` for Resource above to allow the user/group to
Assume Any role. In practice, you might want to automate this as part of the
creation of the relevant roles. (i.e. create the role, then give the relevant
group permissions to Assume that role).
