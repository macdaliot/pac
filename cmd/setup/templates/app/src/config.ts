interface URLConfig {
    apiUrl: string;
    siteUrl: string;
}
let UrlConfig: URLConfig;
if (window.location.host.startsWith("localhost")){
    UrlConfig = { 
        apiUrl: "http://localhost:3000/api/", 
        siteUrl: "http://localhost:8080" 
    };
}
else {
    UrlConfig = { 
        apiUrl: "http://api.foobazbar.pac.pyramidchallenges.com/api/", 
        siteUrl: "http://integration.foobazbar.pac.pyramidchallenges.com"
    };
}

export { UrlConfig };