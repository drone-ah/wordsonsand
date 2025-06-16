---
categories: []
date: "2023-11-08T12:10:16Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:05"
  _last_editor_used_jetpack: block-editor
  _publicize_job_id: "89180754980"
  timeline_notification: "1699445418"
  wordads_ufa: s:wpcom-ufa-v4:1699445541
parent_id: "0"
password: ""
status: publish
tags:
- aws
- aws-lambda
- golang
title: How to get current Function URL (aws-lambda + golang)
type: post
url: /2023/11/08/how-to-get-current-function-url-aws-lambda-golang/
---

When deploying a function lambda that needs details of its own function URL.
It\'s an OAuth Callback, and needs to calculate the redirect. There are possible
security issues doing it this way, so will switch to http gateway on launch. In
the meantime, though, I ran into a bit of a chicken and egg problem.

In Pulumi, the function URL is created after the function (and even otherwise),
I can\'t pass the output of the lambda (or lambdaFunctionUrl) back in as an
environment variable. Fortunately, there is an easy way to pick up the Function
URL (or the function name for that matter) - if you know how ;)

```wp-block-syntaxhighlighter-code
   domainName := request.RequestContext.DomainName
    funcName := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
    return fmt.Sprintf("Domain: %s, funcName: %s", domainName, funcName), nil
```

There are other
[defined lambda function environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime)
as well that you can use.
