import { TestBed } from '@angular/core/testing';

import { DailyEventsService } from './daily-events.service';

describe('DailyEventsService', () => {
  let service: DailyEventsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DailyEventsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
