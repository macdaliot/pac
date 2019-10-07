output "tls" {
    value = tls_private_key.jumpbox
}

output "private_key" {
  value = tls_private_key.jumpbox.private_key_pem
  sensitive = true
}

output "public_key" {
  value = tls_private_key.jumpbox.public_key_pem
  sensitive = true
}

output "public_openssh" {
  value = tls_private_key.jumpbox.public_key_openssh
  sensitive = true
}
