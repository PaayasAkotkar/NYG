import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerFreezeComponent } from './power-freeze.component';

describe('PowerFreezeComponent', () => {
  let component: PowerFreezeComponent;
  let fixture: ComponentFixture<PowerFreezeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerFreezeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerFreezeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
