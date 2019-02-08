interface URLConfig {
    apiUrl: string;
    siteUrl: string;
}
let UrlConfig: URLConfig;
if (window.location.host == "localhost"){
    UrlConfig = { 
        apiUrl: "http://localhost:3000", 
        siteUrl: "http://localhost:8080" 
    };
}
else {
    UrlConfig = { 
        apiUrl: "http://api.{{.projectName}}.pac.pyramidchallenges.com/api", 
        siteUrl: "http://integration.{{.projectName}}.pac.pyramidchallenges.com"
    };
}

export { UrlConfig };