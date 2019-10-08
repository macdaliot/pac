const cloud = {
    region: '{{.region}}'
};

const local = {
    region: 'local',
    endpoint: 'http://pac-{{.projectName}}-db-local:8000'
};

export const awsConfig = { local, cloud };
