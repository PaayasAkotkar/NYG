import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RefRewindPowerComponent } from './ref-rewind-power.component';

describe('RefRewindPowerComponent', () => {
  let component: RefRewindPowerComponent;
  let fixture: ComponentFixture<RefRewindPowerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RefRewindPowerComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RefRewindPowerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
