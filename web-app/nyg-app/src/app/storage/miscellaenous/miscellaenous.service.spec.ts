import { TestBed } from '@angular/core/testing';

import { MiscellaenousService } from './miscellaenous.service';

describe('MiscellaenousService', () => {
  let service: MiscellaenousService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MiscellaenousService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
