provider "acme" {
  server_url = "https://acme-v02.api.letsencrypt.org/directory"
}

resource "acme_registration" "reg" {
  account_key_pem = tls_private_key.jumpbox.private_key_pem
  email_address   = var.acme_registration_email
}

resource "acme_certificate" "certificate" {
  account_key_pem = acme_registration.reg.account_key_pem
  common_name     = "*.${var.project_name}.${var.hosted_zone}"

  dns_challenge {
    provider = "route53"

    config = {
      AWS_HOSTED_ZONE_ID = data.terraform_remote_state.dns.outputs.main_zone_id
    }
  }
}

output "acme_cert" {
  value = acme_certificate.certificate
}