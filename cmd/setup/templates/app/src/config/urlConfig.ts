interface URLConfig {
  apiUrl: string;
  siteUrl: string;
}

let UrlConfig: URLConfig;
if (window.location.host.startsWith('localhost')) {
  UrlConfig = {
    apiUrl: 'http://localhost:3000/api/',
    siteUrl: 'http://localhost:8080'
  };
} else {
  UrlConfig = {
    apiUrl: 'https://api.{{.projectName}}.pac.pyramidchallenges.com/api/',
    siteUrl: 'https://integration.{{.projectName}}.pac.pyramidchallenges.com'
  };
}

export { UrlConfig };
