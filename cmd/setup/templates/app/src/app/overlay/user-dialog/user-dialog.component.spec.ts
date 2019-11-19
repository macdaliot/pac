import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { UserDialogComponent } from './user-dialog.component';
import { ReactiveFormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatMomentDateModule, MAT_MOMENT_DATE_ADAPTER_OPTIONS } from '@angular/material-moment-adapter';
import { MatSelectModule } from '@angular/material/select';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { By } from '@angular/platform-browser';
describe('UserDialogComponent', () => {
 const data = {msg: 'Dialog Message'};
 let component: UserDialogComponent;
 let fixture: ComponentFixture<UserDialogComponent>;

 beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UserDialogComponent ],
      imports: [ ReactiveFormsModule, HttpClientModule,
      MatDatepickerModule, MatInputModule, MatFormFieldModule, MatDialogModule, MatSelectModule, RouterTestingModule, MatMomentDateModule, BrowserAnimationsModule],
      providers: [{ provide: MatDialogRef, useValue: {} }, { provide: MAT_DIALOG_DATA, useValue: data },
      { provide: 'HOST', useValue: window.location.host }
      ]

    })
    .compileComponents();
  }));

 beforeEach(() => {
    fixture = TestBed.createComponent(UserDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

 it('should create', () => {
     expect(component).toBeTruthy();
   });

 it('call save on save click', () => {
    // const button = fixture.debugElement.query(By.css('#saveButton'));
   //  const compiled = fixture.debugElement.nativeElement;
///const butt2 = compiled.querySelector('#saveButton');
   //  debugger; 
  //   expect(button.nativeElement.innerText).toEqual('Save');
    // expect(component).toBeTruthy();
   });

});
