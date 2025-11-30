import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetupDrawPowerComponent } from './setup-draw-power.component';

describe('SetupDrawPowerComponent', () => {
  let component: SetupDrawPowerComponent;
  let fixture: ComponentFixture<SetupDrawPowerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetupDrawPowerComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SetupDrawPowerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
