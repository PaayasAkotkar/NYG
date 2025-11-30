import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerUpPatternComponent } from './power-up-pattern.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('PowerUpPatternComponent', () => {
  let component: PowerUpPatternComponent;
  let fixture: ComponentFixture<PowerUpPatternComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, PowerUpPatternComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(PowerUpPatternComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
