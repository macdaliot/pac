import { Component, OnInit } from '@angular/core';
import { AuthService } from './service/auth/auth.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'ng-app';

  constructor(private authService: AuthService) { }

  ngOnInit() {
  	console.log('App INIT');
  	this.authService.loadUserToken();
  	this.authService.loadUserToken();
  }
}
