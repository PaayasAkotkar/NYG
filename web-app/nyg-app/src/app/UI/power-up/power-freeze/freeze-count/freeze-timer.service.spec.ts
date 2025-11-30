import { TestBed } from '@angular/core/testing';

import { FreezeTimerService } from './freeze-timer.service';

describe('FreezeTimerService', () => {
  let service: FreezeTimerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FreezeTimerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
