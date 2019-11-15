import { TestBed } from '@angular/core/testing';

import { AuthService } from './auth.service';

import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientModule } from '@angular/common/http';
describe('AuthService', () => {
  beforeEach(() => TestBed.configureTestingModule({
  	imports: [RouterTestingModule, HttpClientModule],
  	providers: [{ provide: 'HOST', useValue: window.location.host }]

  }));
  // Sunday, March 15, 2015 12:00:00 PM
  // Wednesday, December 31, 3000 11:59:59 PM
  it('should be created', () => {
    const service: AuthService = TestBed.get(AuthService);
    expect(service).toBeTruthy();
  });

  it('should not be logged in', () => {
  	const service: AuthService = TestBed.get(AuthService);
   expect(service.isAuthenticated()).toBe(false);
  });

  it('should not accept expired token', () => {
  	const service: AuthService = TestBed.get(AuthService);
  	const jwtToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkpXVCBUZXN0ZXIiLCJleHAiOjE0MjY0MjA4MDB9.IGw2tfQ-JfLo_lPxIzAUZLd54alQWsMOr0kydxczmKg';

  	expect(service.login(jwtToken)).toBe(false);
  	expect(service.isAuthenticated()).toBe(false);
  });



  it('should accept unexpired token', () => {
  	const service: AuthService = TestBed.get(AuthService);
  	const jwtToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkpXVCBUZXN0ZXIiLCJleHAiOjMyNTM1MjE1OTk5fQ.t8-Micm5xDAwGrl-3XSsY74xpm0N7cAbus30_nnz_NE';

  	expect(service.login(jwtToken)).toBe(true);
  	expect(service.isAuthenticated()).toBe(true);
  });

  it('should log-in and then log-out', () => {
  	const service: AuthService = TestBed.get(AuthService);
  	const jwtToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkpXVCBUZXN0ZXIiLCJleHAiOjMyNTM1MjE1OTk5fQ.t8-Micm5xDAwGrl-3XSsY74xpm0N7cAbus30_nnz_NE';

  	expect(service.login(jwtToken)).toBe(true);
  	expect(service.isAuthenticated()).toBe(true);
   expect(service.canActivate()).toBe(true);

  	service.logout();
	  expect(service.isAuthenticated()).toBe(false);
   expect(service.canActivate()).toBe(false);
  });
});
