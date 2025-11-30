import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerRewindComponent } from './power-rewind.component';

describe('PowerRewindComponent', () => {
  let component: PowerRewindComponent;
  let fixture: ComponentFixture<PowerRewindComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerRewindComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerRewindComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
