const cloud = {
    region: 'us-east-2'
};

const local = {
    region: 'local',
    endpoint: 'http://pac-db-local:8000'
};

export const awsConfig = { local, cloud };