output "tls" {
    value = tls_private_key.management
}

output "private_key" {
  value = tls_private_key.management.private_key_pem
  sensitive = true
}

output "public_key" {
  value = tls_private_key.management.public_key_pem
  sensitive = true
}

output "public_openssh" {
  value = tls_private_key.management.public_key_openssh
  sensitive = true
}
