import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from '../../service/auth/auth.service';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(private route: ActivatedRoute, private router: Router, private authService: AuthService) { }

  ngOnInit() {
    console.log('LOGIN INIT');
    this.route.queryParams.subscribe(params => {
    console.log('SUBSCRIBE PARAMS');

    if (params.jwt) {
        this.authService.login(params.jwt);
        this.router.navigate(['home']);
      }
    });
  }

}
