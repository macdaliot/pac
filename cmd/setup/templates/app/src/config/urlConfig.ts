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
    apiUrl: 'http://api.{{.projectName}}.pac.pyramidchallenges.com/api/',
    siteUrl:
      'http://integration.{{.projectName}}.pac.pyramidchallenges.com.s3-website.us-east-2.amazonaws.com'
  };
}

export { UrlConfig };
