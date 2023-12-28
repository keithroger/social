resource "aws_api_gateway_domain_name" "this" {
  certificate_arn = var.certificate_arn
  domain_name     = var.domain
}

resource "aws_route53_record" "this" {
  name    = aws_api_gateway_domain_name.this.domain_name
  type    = "A"
  zone_id = var.route_53_hosted_zone_id

  alias {
    evaluate_target_health = true
    name                   = aws_api_gateway_domain_name.this.cloudfront_domain_name
    zone_id                = aws_api_gateway_domain_name.this.cloudfront_zone_id
  }
}