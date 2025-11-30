import { TestBed } from '@angular/core/testing';

import { SetScoreboardService } from './set-scoreboard.service';

describe('SetScoreboardService', () => {
  let service: SetScoreboardService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SetScoreboardService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
