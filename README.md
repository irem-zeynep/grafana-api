# Grafana API

Exercise project for DevOps position.
Includes simple AWS API Gateway, Lambda Function which interacts with original Grafana client, SNS, timeStream and secret manager.

## Required Secrets

Grafana Client
```json
{
  "apiKey": "<grafana-api-key>",
  "host": "<grafana-host>"
}
```

SNS
```json
{
  "topicName": "<sns-error-topic-name>"
}
```

TimeStream Table
```json
{
  "dbName": "<terraform-db-name>",
  "tableName": "<terraform-table-name>"
}
```

## Required Variables for Terraform
```shell
export TF_VAR_aws_access_key=<access_key_value>
export TF_VAR_aws_secret_key=<secret_key_value>
export TF_VAR_aws_region=<region>         
```

## Enable Pre-commit Hook

Install pre-commit
``` shell
pip install pre-commit
```

Install linter for golang (for mac)
``` shell
brew install golangci-lint
```

Enable pre-commit
``` shell
make precommit-install
```

## Create Mock files for unit testing
``` shell
make generate-mocks
```

## Tips

- Installed and run grafana instance on ec2 instance
- Used SQS to monitor SNS events