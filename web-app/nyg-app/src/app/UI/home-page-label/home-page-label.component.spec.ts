import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePageLabelComponent } from './home-page-label.component';

describe('HomePageLabelComponent', () => {
  let component: HomePageLabelComponent;
  let fixture: ComponentFixture<HomePageLabelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HomePageLabelComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomePageLabelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
