import { Route } from '@angular/router';
import { HomeComponent } from './views/home/home.component';
import { LoginComponent } from './views/login/login.component';
import { PageOneComponent } from './views/page-one/page-one.component';
import { PageTwoComponent } from './views/page-two/page-two.component';
import { AuthService } from './service/auth/auth.service';
export interface NamedRoute extends Route {
	displayName?: string;
}

export const routes: Array<NamedRoute> = [
		{ path: 'home', component: HomeComponent, displayName: 'Home'},
		{ path: 'page1', component: PageOneComponent },
		{ path: 'page2', component: PageTwoComponent, displayName: 'Page Two'},
		{ path: 'login', component: LoginComponent },
		{ path: '', redirectTo: '/home', pathMatch: 'full' }
];

// , canActivate: [AuthService]
