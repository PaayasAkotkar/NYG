import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WatcherTimerComponent } from './timer.component';

describe('WatcherTimerComponent', () => {
  let component: WatcherTimerComponent;
  let fixture: ComponentFixture<WatcherTimerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WatcherTimerComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(WatcherTimerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
