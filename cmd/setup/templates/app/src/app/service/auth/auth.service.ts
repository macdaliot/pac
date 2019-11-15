import { Injectable } from '@angular/core';
import { IUser } from '@pyramid-systems/domain';
import { BehaviorSubject, Observable } from 'rxjs';
import { JWT_TOKEN_NAME } from '../../constants';
import { Router, CanActivate } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService implements CanActivate {

constructor(private router: Router) {}

  private _currentUser = new BehaviorSubject(null);

  private _userToken = null;

  get currentUser(): Observable<any> {
    return this._currentUser.asObservable();
  }

  private get userToken() {
  	return this._userToken;
  }

  private set userToken(val) {
  	this._userToken = val;
  	this._currentUser.next(val);
  }

  private getJsonToken(token: string): string {
     const payload = token.split('.')[1]; // lop off header and signature
     const json = window.atob(payload).toString();
     return JSON.parse(json);
  }

  login(token: string) {
  	const userToken = this.getJsonToken(token);
   if (this.isAuthenticated(userToken)) {
      this.userToken = userToken;
      localStorage.setItem(JWT_TOKEN_NAME, token);
      return true;
    }
   return false;
  }

  logout() {
  	this.userToken = null;
  	localStorage.removeItem(JWT_TOKEN_NAME);
  }

  isAuthenticated(token = null): boolean {
  	if (!token) {
  		token = this.userToken;
  	}

  	if (token) {
  		return ((token.exp * 1000) > Date.now());
  	}
 	 return false;
  }

  loadUserToken() {
    if (!this.userToken) {
    	const encodedToken = localStorage.getItem(JWT_TOKEN_NAME);
     if (encodedToken) {
      // 	const jsonToken = this.getJsonToken(encodedToken);
        this.login(encodedToken);
      	// if (encodedToken && this.isAuthenticated(jsonToken)) {
      	// 	this.login(encodedToken);
       //    return;
      	// }
      }
    }
  }

  canActivate(): boolean {
    if (!this.isAuthenticated()) {
      this.logout();
      this.router.navigate(['home']);
      return false;
    }
    return true;
  }
}



