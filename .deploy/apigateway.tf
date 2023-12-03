resource "aws_api_gateway_rest_api" "grafana_api" {
  name        = "grafana-api"
  description = "Grafana API Gateway"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_resource" "root" {
  rest_api_id = aws_api_gateway_rest_api.grafana_api.id
  parent_id   = aws_api_gateway_rest_api.grafana_api.root_resource_id
  path_part   = "organization-users"
}

resource "aws_api_gateway_method" "endpoint" {
  rest_api_id        = aws_api_gateway_rest_api.grafana_api.id
  resource_id        = aws_api_gateway_resource.root.id
  http_method        = "GET"
  authorization      = "NONE"
  request_parameters = {
    "method.request.querystring.org"   = true,
    "method.request.querystring.email" = true
  }
}

resource "aws_api_gateway_integration" "lambda_integration" {
  rest_api_id             = aws_api_gateway_rest_api.grafana_api.id
  resource_id             = aws_api_gateway_resource.root.id
  http_method             = aws_api_gateway_method.endpoint.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda.invoke_arn
}

resource "aws_api_gateway_method_response" "endpoint" {
  rest_api_id = aws_api_gateway_rest_api.grafana_api.id
  resource_id = aws_api_gateway_resource.root.id
  http_method = aws_api_gateway_method.endpoint.http_method
  status_code = "200"
}

resource "aws_api_gateway_deployment" "deployment" {
  depends_on = [
    aws_api_gateway_integration.lambda_integration
  ]

  rest_api_id = aws_api_gateway_rest_api.grafana_api.id
  stage_name  = "stage"
}

output "base_url" {
  value = aws_api_gateway_deployment.deployment.invoke_url
}