# Configure DNS record in aws
data "aws_route53_zone" "selected" {
  name         = "metadlw.com"
}

resource "aws_route53_record" "api" {
  zone_id = data.aws_route53_zone.selected.zone_id
  name    = var.backendDNS
  type    = "A"
  ttl     = 300
  records = [azurerm_public_ip.gwIp.ip_address]
}

resource "aws_route53_record" "web" {
  zone_id = data.aws_route53_zone.selected.zone_id
  name    = var.frontendDNS
  type    = "CNAME"
  ttl     = 10

  records = [ azurerm_cdn_endpoint.dlw_origin.fqdn ]

  depends_on = [
    azurerm_cdn_endpoint.dlw_origin,
    azurerm_cdn_endpoint_custom_domain.dlw_dns
  ]
}
