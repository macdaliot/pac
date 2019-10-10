resource "acme_registration" "reg" {
  account_key_pem = data.terraform_remote_state.bootstrap.outputs.worker_private_key
  email_address   = "jdiederiks@pyramidsystems.com"
}

resource "acme_certificate" "certificate" {
  account_key_pem = "${acme_registration.reg.account_key_pem}"
  common_name     = "*.${var.project_name}.${var.hosted_zone}"

  dns_challenge {
    provider = "route53"

    config = {
      AWS_HOSTED_ZONE_ID = "${data.terraform_remote_state.dns.outputs.main_zone_id}"
    }
  }
}

output "acme_cert" {
  value = acme_certificate.certificate
}
