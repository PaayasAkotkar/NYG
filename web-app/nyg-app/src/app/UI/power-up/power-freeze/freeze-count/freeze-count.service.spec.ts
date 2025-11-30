import { TestBed } from '@angular/core/testing';

import { FreezeCountService } from './freeze-count.service';

describe('FreezeCountService', () => {
  let service: FreezeCountService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FreezeCountService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
