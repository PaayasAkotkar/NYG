import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PhysicsClockComponent } from './physics-clock.component';

describe('PhysicsClockComponent', () => {
  let component: PhysicsClockComponent;
  let fixture: ComponentFixture<PhysicsClockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PhysicsClockComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PhysicsClockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
