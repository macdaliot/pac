const cloud = {
    region: '[psi[.region]]'
};

const local = {
    region: 'local',
    endpoint: 'http://pac-[psi[.projectName]]-db-local:8000'
};

export const awsConfig = { local, cloud };
