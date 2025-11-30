import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerUpComponent } from './power-up.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('PowerUpComponent', () => {
  let component: PowerUpComponent;
  let fixture: ComponentFixture<PowerUpComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerUpComponent, ReactiveFormsModule]
    })
      .compileComponents();

    fixture = TestBed.createComponent(PowerUpComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
