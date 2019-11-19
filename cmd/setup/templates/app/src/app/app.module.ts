import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { GovWebsiteBannerComponent } from './components/gov-website-banner/gov-website-banner.component';
import { HeaderComponent } from './components/header/header.component';
import { LoginCallbackComponent } from './components/login-callback/login-callback.component';
import { NavigationBarComponent } from './components/navigation-bar/navigation-bar.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { HomeComponent } from './views/home/home.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatTableModule } from '@angular/material/table';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatDatepickerModule } from '@angular/material/datepicker';
// import { MatNativeDateModule } from '@angular/material/core'
import { MatMomentDateModule, MAT_MOMENT_DATE_ADAPTER_OPTIONS } from '@angular/material-moment-adapter';
import { MatSelectModule } from '@angular/material/select';

import { PageOneComponent } from './views/page-one/page-one.component';
import { PageTwoComponent } from './views/page-two/page-two.component';
import { AlertCardComponent } from './components/alert-card/alert-card.component';
import { LoginComponent } from './views/login/login.component';
import { PhonePipe } from './pipe/phone.pipe';
import { UserDialogComponent } from './overlay/user-dialog/user-dialog.component';
/// import { ConfirmDeleteDialogComponent } from './confirm-delete-dialog/confirm-delete-dialog.component';
@NgModule({
  declarations: [
    AppComponent,
    GovWebsiteBannerComponent,
    HeaderComponent,
    LoginCallbackComponent,
    NavigationBarComponent,
    HomeComponent,
    PageOneComponent,
    PageTwoComponent,
    AlertCardComponent,
    LoginComponent,
    PhonePipe,
    UserDialogComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FontAwesomeModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatButtonModule,
    MatCardModule,
    MatButtonToggleModule,
    MatTableModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
   // MatNativeDateModule,
    MatMomentDateModule,
    MatSelectModule
  ],
  entryComponents: [UserDialogComponent],
  providers: [{ provide: 'HOST', useValue: window.location.host },
              {provide: MAT_MOMENT_DATE_ADAPTER_OPTIONS, useValue: {useUtc: true}},
              MatDatepickerModule],
  bootstrap: [AppComponent]
})
export class AppModule { }
