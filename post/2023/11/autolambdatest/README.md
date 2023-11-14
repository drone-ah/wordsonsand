# Introduction

Companion code for the [blog post about automated testing of AWS Lambda locally](https://drone-ah.com/2023/11/14/automated-testing-of-aws-lambda-locally/)

# Prerequisites

* Bazel - though technically `go test` might work, but I've not tested it
* Pulumi - you'll need to have logged in from the console
* AWS Account (it will use the currently active profile to deploy a couple of resources)
* AssumeRole permission for current user (Assume your user has the `AdministratorAccess` policy)

# Running

You can run the test easily with:

```shell
bazel test ... --define="gotags=integration_test" --sandbox_writable_path=$HOME/.pulumi
```

the gotags bit allows [separating out integration tests from unit tests](https://drone-ah.com/2023/11/13/separating-out-integration-tests-for-golang-in-bazel/)

The `sandbox_writable_path` flag [circumvents an issue around pulumi wanting to write to `~/.credentials.json`](https://drone-ah.com/2023/11/14/unable-to-write-to-home-pulumi-credentials-json-during-bazel-test/)
