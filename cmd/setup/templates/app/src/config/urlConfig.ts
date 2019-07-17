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
    apiUrl: 'https://' + location.host + '/api/',
    siteUrl: 'https://' + location.host
  };
}

export { UrlConfig };
