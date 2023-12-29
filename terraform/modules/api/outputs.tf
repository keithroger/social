output "api_gateway_uri" {
  value = aws_apigatewayv2_api.this.api_endpoint
}

output "vpc_link_sg_id" {
  value = aws_security_group.vpc_link.id
}