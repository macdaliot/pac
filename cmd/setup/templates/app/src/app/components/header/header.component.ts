import { Component, OnInit } from '@angular/core';
import { faUser, faCaretDown } from '@fortawesome/free-solid-svg-icons';
import { UrlService } from '../../service/url/url.service';
import { AuthService } from '../../service/auth/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  username: string;
  expiry: number;
  faUser = faUser;
  faCaretDown = faCaretDown;
  constructor(private urlService: UrlService, private authService: AuthService) {
    this.authService.currentUser.subscribe(res => {
      this.username = res && res.name;
      this.expiry = res && (res.exp * 1000);
    });
  }

  ngOnInit() {

  }

  gotoLogin() {
  	window.open(this.urlService.apiUrl + 'auth/login', '_self');
  }

  logout() {
    this.authService.logout();
  }


}
