import { Component, OnInit } from '@angular/core';
import { NamedRoute, routes } from '../../routes';

@Component({
  selector: 'app-navigation-bar',
  templateUrl: './navigation-bar.component.html',
  styleUrls: ['./navigation-bar.component.scss']
})
export class NavigationBarComponent implements OnInit {
  displayRoutes: Array<NamedRoute>;
  constructor() {
  		this.displayRoutes = routes.filter(r => r.displayName);
  }

  ngOnInit() {
  }

}
