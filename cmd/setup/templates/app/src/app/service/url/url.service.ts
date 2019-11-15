import { Injectable, Inject } from '@angular/core';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class UrlService {
  apiUrl: string;
  siteUrl: string;

  constructor(private router: Router, @Inject('HOST') private host: string) {
 	  const hostStr = host;

   	if (hostStr.indexOf('localhost') === 0) {
   		this.apiUrl = 'http://localhost:3000/api/';
   		this.siteUrl = 'http://' + hostStr;
   	} else {
  	  this.apiUrl = 'https://' + hostStr + '/api/';
     this.siteUrl = 'https://' + hostStr;
   	}
   	console.log(this.apiUrl, this.siteUrl);
  }

  get hostname() {
  	return this.host;
  }
}

