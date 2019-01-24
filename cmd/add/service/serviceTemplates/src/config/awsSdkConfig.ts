const cloud = {
    region: 'us-east-2'
};

const local = {
    region: 'local',
    endpoint: 'http://pac-db-local:8000'
};

const awsConfig = { "local" : local, "cloud" : cloud };
export default awsConfig;