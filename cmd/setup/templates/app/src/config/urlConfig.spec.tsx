import { UrlConfig } from '.';

describe('urlconfig', () => {
    it("should return urlconfig for local env", () => {
        expect(UrlConfig.apiUrl).toBe("http://localhost:3000/api/");
        expect(UrlConfig.siteUrl).toBe("http://localhost:8080")
    })
});
