import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PageTwoComponent } from './page-two.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { MatTableModule } from '@angular/material/table';
import { PhonePipe } from '../../pipe/phone.pipe';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientModule } from '@angular/common/http';
import { MatDialogModule } from '@angular/material/dialog';

describe('PageTwoComponent', () => {
  let component: PageTwoComponent;
  let fixture: ComponentFixture<PageTwoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PageTwoComponent, PhonePipe ],
      imports: [ FontAwesomeModule, MatTableModule, RouterTestingModule, MatDialogModule, HttpClientModule ],
       providers: [{ provide: 'HOST', useValue: window.location.host }]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PageTwoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
