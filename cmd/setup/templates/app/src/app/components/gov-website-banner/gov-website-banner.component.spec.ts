import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GovWebsiteBannerComponent } from './gov-website-banner.component';

describe('GovWebsiteBannerComponent', () => {
  let component: GovWebsiteBannerComponent;
  let fixture: ComponentFixture<GovWebsiteBannerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GovWebsiteBannerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GovWebsiteBannerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
