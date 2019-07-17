data "aws_kms_key" "project_key" {
  key_id = "{{ .encryptionKeyID }}"
}
