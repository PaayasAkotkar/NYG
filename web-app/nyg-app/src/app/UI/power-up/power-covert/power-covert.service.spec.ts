import { TestBed } from '@angular/core/testing';

import { PowerCovertService } from './power-covert.service';

describe('PowerCovertService', () => {
  let service: PowerCovertService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PowerCovertService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
