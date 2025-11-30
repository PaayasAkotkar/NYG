import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerDrawComponent } from './power-draw.component';

describe('PowerDrawComponent', () => {
  let component: PowerDrawComponent;
  let fixture: ComponentFixture<PowerDrawComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerDrawComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerDrawComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
