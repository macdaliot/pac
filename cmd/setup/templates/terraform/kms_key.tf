data "aws_kms_key" "{{ .projectName }}_kms_key" {
  key_id = "{{ .encryptionKeyID }}"
}