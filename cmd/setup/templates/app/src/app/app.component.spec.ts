import { TestBed, async } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';

import { PageOneComponent } from './views/page-one/page-one.component';
import { PageTwoComponent } from './views/page-two/page-two.component';
import { AlertCardComponent } from './components/alert-card/alert-card.component';
import { LoginComponent } from './views/login/login.component';
import { PhonePipe } from './pipe/phone.pipe';
import { UserDialogComponent } from './overlay/user-dialog/user-dialog.component';
import { GovWebsiteBannerComponent } from './components/gov-website-banner/gov-website-banner.component';
import { HeaderComponent } from './components/header/header.component';
import { NavigationBarComponent } from './components/navigation-bar/navigation-bar.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

describe('AppComponent', () => {
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        RouterTestingModule, FontAwesomeModule
      ],
      declarations: [
        AppComponent, GovWebsiteBannerComponent, HeaderComponent, NavigationBarComponent
      ],
      providers: [{ provide: 'HOST', useValue: window.location.host }]
    }).compileComponents();
  }));

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have as title 'ng-app'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('ng-app');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1.application-title').textContent).toContain('[psi[.projectName]]');
  });
});
