import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { AlertCardComponent } from '../../components/alert-card/alert-card.component';
import { PageOneComponent } from './page-one.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
describe('PageOneComponent', () => {
  let component: PageOneComponent;
  let fixture: ComponentFixture<PageOneComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PageOneComponent, AlertCardComponent ],
      imports: [FontAwesomeModule]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PageOneComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
