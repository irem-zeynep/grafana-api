# Grafana API

Exercise project for DevOps position.
Includes simple AWS API Gateway, Lambda Function which interacts with original Grafana client, SNS, timeStream and secret manager.

## Required Secrets

Grafana Client (`grafana/stage/grafanaclient`)
```json
{
  "apiKey": "<grafana-api-key>",
  "host": "<grafana-host>"
}
```

SNS (`grafana/stage/errortopic`)
```json
{
  "topicName": "<sns-error-topic-name>"
}
```

TimeStream Table (`grafana/stage/timestreamdb`)
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

## Additional Tips

- Installed and run Grafana instance on ec2 instance
- Created SNS topic and timeStream DB
- Used SQS to monitor SNS events
- Monitored logs with CloudWatch

### Enable Pre-commit Hook

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

### Create mocks for unit testing
``` shell
make generate-mocks
```
