# Overview
Study of [vektra/mockery](https://vektra.github.io/mockery/latest/) usage.


# 1. How to create mock
```
% docker pull vektra/mockery
% docker run -v "$PWD":/src -w /src vektra/mockery --all
```

# 2. How to test
see : [./service/user_service_test.go](./service/user_service_test.go)