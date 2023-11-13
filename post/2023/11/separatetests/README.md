# Introduction

Contains the source files / samples for the blog post about [Separating out integration tests for golang in Bazel](https://drone-ah.com/2023/11/13/separating-out-integration-tests-for-golang-in-bazel/)

To (re)build the BUILD files including unit & integration tests
```shell
bazel run //:gazelle -- -build_tags=unit_test,integration_test
```

You can also just `bazel run //:gazelle` in this repo because the `BUILD` file switches these two tags on with:
```
# gazelle:build_tags unit_test,integration_test
```

To run unit tests
```shell
bazel test ... --define gotags=unit_test
```

To run integration tests:
```shell
bazel test ... --define gotags=integration_test
```

To run both:
```shell
bazel test ... --define gotags=unit_test,integration_test
```
