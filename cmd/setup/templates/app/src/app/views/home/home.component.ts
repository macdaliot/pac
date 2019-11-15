import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { AuthService } from '../../service/auth/auth.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  username: string;

  constructor(private route: ActivatedRoute, private authService: AuthService) { }

   ngOnInit() {
    this.authService.currentUser.subscribe(res => {
      this.username = res && res.name;
    });
 	}

}
