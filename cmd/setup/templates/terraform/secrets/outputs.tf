output "jwt_issuer" {
  value = aws_ssm_parameter.jwt_issuer
}

output "jwt_secret" {
  value = aws_ssm_parameter.jwt_secret
}
