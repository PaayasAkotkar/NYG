import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GymClockComponent } from './gym-clock.component';

describe('GymClockComponent', () => {
  let component: GymClockComponent;
  let fixture: ComponentFixture<GymClockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GymClockComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GymClockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
