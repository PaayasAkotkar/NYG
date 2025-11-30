import { TestBed } from '@angular/core/testing';

import { CheatService } from './cheat.service';

describe('CheatService', () => {
  let service: CheatService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CheatService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
