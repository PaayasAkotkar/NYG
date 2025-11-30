import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LocifyClockComponent } from './locify-clock.component';

describe('LocifyClockComponent', () => {
  let component: LocifyClockComponent;
  let fixture: ComponentFixture<LocifyClockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LocifyClockComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LocifyClockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
