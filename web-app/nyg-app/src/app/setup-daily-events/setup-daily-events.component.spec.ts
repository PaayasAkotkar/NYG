import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetupDailyEventsComponent } from './setup-daily-events.component';

describe('SetupDailyEventsComponent', () => {
  let component: SetupDailyEventsComponent;
  let fixture: ComponentFixture<SetupDailyEventsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetupDailyEventsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SetupDailyEventsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
