interface Environment {
  endpoint?: string,
  region: string
}

interface AwsConfig {
  [name: string]: Environment
}

const cloud: Environment = {
    region: '{{ .region }}'
};

const local: Environment = {
    region: 'local',
    endpoint: 'http://pac-{{.projectName}}-db-local:8000'
};

export const awsConfig: AwsConfig = {
  cloud,
  local
};