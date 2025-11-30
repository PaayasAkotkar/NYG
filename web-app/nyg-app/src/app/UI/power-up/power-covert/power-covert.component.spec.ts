import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerCovertComponent } from './power-covert.component';

describe('PowerCovertComponent', () => {
  let component: PowerCovertComponent;
  let fixture: ComponentFixture<PowerCovertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerCovertComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerCovertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
