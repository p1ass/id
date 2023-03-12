# terraform

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/docs)
- [Terragrunt](https://terragrunt.gruntwork.io/)
- [direnv](https://github.com/direnv/direnv)

You can install dependencies:

```shell
make deps
```

## Setup

```shell
cp .env.template .env
# Edit environment values
vim .env
direnv allow
```

## Format

```shell
make fmt
```

## Exec on your laptop

```shell
gcloud auth application-default login
make plan
```
