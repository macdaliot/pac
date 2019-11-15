import { TestBed } from '@angular/core/testing';

import { UrlService } from './url.service';

import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientModule } from '@angular/common/http';

describe('UrlService non-localhost', () => {
  beforeEach(() => TestBed.configureTestingModule({
  	imports: [RouterTestingModule, HttpClientModule],
  	providers: [{ provide: 'HOST', useValue: 'test' }]

  }));

  it('should be created', () => {
    const service: UrlService = TestBed.get(UrlService);
    expect(service).toBeTruthy();
  });

  it('determines URLs', () => {
    const service: UrlService = TestBed.get(UrlService);

  	 expect(service.apiUrl).toBe('https://test/api/');
 	  expect(service.siteUrl).toBe('https://test');

  });
});


describe('UrlService localhost', () => {
  beforeEach(() => TestBed.configureTestingModule({
  	imports: [RouterTestingModule, HttpClientModule],
  	providers: [{ provide: 'HOST', useValue: 'localhost' }]

  }));

  it('should be created', () => {
    const service: UrlService = TestBed.get(UrlService);
    expect(service).toBeTruthy();
  });

  it('determines URLs', () => {
    const service: UrlService = TestBed.get(UrlService);

  	 expect(service.apiUrl).toBe('http://localhost:3000/api/');
 	  expect(service.siteUrl).toBe('http://localhost');

  });

  it('echo hostname', () => {
  	 const service: UrlService = TestBed.get(UrlService);
  	 expect(service.hostname).toBe('localhost');
  });
});
