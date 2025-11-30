import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeProfileHeaderComponent } from './home-profile-header.component';

describe('HomeProfileHeaderComponent', () => {
  let component: HomeProfileHeaderComponent;
  let fixture: ComponentFixture<HomeProfileHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HomeProfileHeaderComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeProfileHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
