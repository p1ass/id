# backend

Backend system for https://id.p1ass.com written in Go

## Requirements

- Go 1.19 or later

## Getting Started

1. Execute application

```shell
go run main.go
```

## Build Image

### Pre requirements

Install [ko](https://github.com/google/ko)

```shell
 go install github.com/google/ko@latest
```

### Build Image for local

```shell
ko build --local --base-import-paths .

# or you can run container locally:
docker run -p 8080:8080 $(ko build --local --base-import-paths .)
```

### Push to Google Artifact Registry
```shell
cp .env.template .env
# Edit environment values
vim .env
ko build --base-import-paths . --platform=linux/amd64,linux/arm64 
```