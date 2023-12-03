data "archive_file" "lambda_package" {
  type        = "zip"
  source_file = "../cmd/lambda/check-organization-user/main"
  output_path = "./main.zip"
}

resource "aws_lambda_function" "lambda" {
  filename         = data.archive_file.lambda_package.output_path
  function_name    = "checkOrganizationUser"
  role             = "arn:aws:iam::786099939343:role/GrafanaOrganizationUserLambdaRole"
  handler          = "main"
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda_package.output_base64sha256
  ephemeral_storage { size = 512 }
  memory_size      = 512
  timeout          = 15
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.grafana_api.execution_arn}/*/GET/${aws_api_gateway_resource.root.path_part}"
}