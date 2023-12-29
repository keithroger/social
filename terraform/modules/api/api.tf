resource "aws_apigatewayv2_api" "this" {
  name          = var.name
  protocol_type = "HTTP"

  tags = {
    name = "${var.name}"
  }
}


# resource "aws_apigatewayv2_integration" "vpc_link" {
#   api_id = aws_apigatewayv2_api.vpc_link.id
#   #   credentials_arn  = aws_iam_role.example.arn
#   description      = "Request to discovery service"
#   integration_type = "HTTP_PROXY"
#   integration_uri  = aws_service_discovery_service.service_discovery.arn

#   integration_method = "ANY"
#   connection_type    = "VPC_LINK"
#   connection_id      = aws_apigatewayv2_vpc_link.vpc_link.id
# }

# resource "aws_apigatewayv2_route" "vpc_link" {
#   api_id    = aws_apigatewayv2_api.vpc_link.id
#   route_key = "ANY /example/{proxy+}"

#   target = "integrations/${aws_apigatewayv2_integration.vpc_link.id}"
# }

resource "aws_apigatewayv2_route" "ecs_nlb" {
  api_id    = aws_apigatewayv2_api.this.id
  route_key = "ANY /test"

  target = "integrations/${aws_apigatewayv2_integration.ecs_nlb.id}"
}

resource "aws_apigatewayv2_route" "test_route" {
  api_id    = aws_apigatewayv2_api.this.id
  route_key = "ANY /"

  target = "integrations/${aws_apigatewayv2_integration.ecs_nlb.id}"
}

resource "aws_apigatewayv2_integration" "ecs_nlb" {
  api_id = aws_apigatewayv2_api.this.id
  #   credentials_arn  = aws_iam_role.example.arn
  description      = "Send requests to ECS load balancer"
  integration_type = "HTTP_PROXY"
  integration_uri  = var.listener_arn

  integration_method = "ANY"
  connection_type    = "VPC_LINK"
  connection_id      = aws_apigatewayv2_vpc_link.vpc_link.id
}

resource "aws_apigatewayv2_stage" "dev_stage" {
  api_id      = aws_apigatewayv2_api.this.id
  name        = "dev"
  auto_deploy = true
}

resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.this.id
  name        = "$default"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.access_logs.arn
    format = jsonencode({
      requestId         = "$context.requestId",
      extendedRequestId = "$context.extendedRequestId"
      ip                = "$context.identity.sourceIp"
      caller            = "$context.identity.caller"
      user              = "$context.identity.user"
      requestTime       = "$context.requestTime"
      httpMethod        = "$context.httpMethod"
      resourcePath      = "$context.resourcePath"
      status            = "$context.status"
      protocol          = "$context.protocol"
      responseLength    = "$context.responseLength"
      }
    )
  }
}

resource "aws_cloudwatch_log_group" "access_logs" {
  name = "${var.name}-access-log-group"
}
