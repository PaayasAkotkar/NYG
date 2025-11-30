import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePageProfilesComponent } from './home-page-profiles.component';

describe('HomePageProfilesComponent', () => {
  let component: HomePageProfilesComponent;
  let fixture: ComponentFixture<HomePageProfilesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HomePageProfilesComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomePageProfilesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
