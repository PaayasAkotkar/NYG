import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerSetttingsPatternComponent } from './power-setttings-pattern.component';

describe('PowerSetttingsPatternComponent', () => {
  let component: PowerSetttingsPatternComponent;
  let fixture: ComponentFixture<PowerSetttingsPatternComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerSetttingsPatternComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerSetttingsPatternComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
