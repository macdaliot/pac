output "name" {
  value = aws_key_pair.key.key_name
}

output "public_key" {
  value = tls_private_key.key.public_key_pem
}

output "public_openssh" {
  value = tls_private_key.key.public_key_openssh
}

output "private_key" {
  value = tls_private_key.key.private_key_pem
}
