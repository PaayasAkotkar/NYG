import { TestBed } from '@angular/core/testing';

import { SportsService } from './setup-sports.service';

describe('SetupMenuService', () => {
  let service: SportsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SportsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
