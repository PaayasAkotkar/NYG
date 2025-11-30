import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerBetComponent } from './power-bet.component';

describe('PowerBetComponent', () => {
  let component: PowerBetComponent;
  let fixture: ComponentFixture<PowerBetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerBetComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerBetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
