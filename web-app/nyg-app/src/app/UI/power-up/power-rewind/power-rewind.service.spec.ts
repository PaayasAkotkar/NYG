import { TestBed } from '@angular/core/testing';

import { PowerRewindService } from './power-rewind.service';

describe('PowerRewindService', () => {
  let service: PowerRewindService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PowerRewindService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
