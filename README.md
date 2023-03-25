# jwt

The JWT package provides a simple and efficient way to encode, sign, and verify JSON Web Tokens in Go applications. JSON Web Tokens are a popular standard for representing claims securely between two parties, such as a client and a server, and are widely used in modern web applications and APIs.

## Installation

```bash
go get -u github.com/ginger-go/jwt
```

## Documentation

Please refer to [https://ginger-go.gitbook.io/jwt/](https://ginger-go.gitbook.io/jwt/) for the full documentation.

## Perform tests

Run ginkgo tests with the following command:
```bash
ginkgo -r -v -p --cover --coverpkg=github.com/ginger-go/jwt
```

Read the coverage report with the following command:
```bash
go tool cover -html=coverprofile.out
```
