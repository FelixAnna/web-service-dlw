# Configure the AWS Provider
resource "aws_route53_zone" "main" {
  name = "metadlw.com"
}

resource "aws_route53_record" "dev-api" {
  zone_id = aws_route53_zone.main.zone_id
  name    = "api.metadlw.com"
  type    = "A"
  ttl     = 300
  records = [azurerm_public_ip.gwIp.ip_address]
}
