application_cidr_block = "{{ .vpcCidrBlock }}"
enable_elasticsearch    = "true"
enable_jumpbox          = "true"
enable_documentdb       = "true"
documentdb_user         = "pyramid" #need to move this out of here
documentdb_password     = "password" #need to move this out of here
end_user_cidr           = "{{.endUserIP}}/32"
environment_abbr        = "{{.environmentAbbr}}"
environment_name        = "{{.environmentName}}"
es_instance_type        = "r4.large.elasticsearch"
hosted_zone             = "{{.hostedZone}}"
project_fqdn            = "{{.projectFqdn}}"
project_name            = "{{.projectName}}"
region                  = "{{.region}}"
tracing_active          = "Active"

