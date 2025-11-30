import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePageNav2Component } from './home-page-nav2.component';

describe('HomePageNav2Component', () => {
  let component: HomePageNav2Component;
  let fixture: ComponentFixture<HomePageNav2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HomePageNav2Component]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomePageNav2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
