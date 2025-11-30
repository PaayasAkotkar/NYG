import { TestBed } from '@angular/core/testing';

import { ClockTimerService } from './clock-timer.service';

describe('ClockTimerService', () => {
  let service: ClockTimerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ClockTimerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
