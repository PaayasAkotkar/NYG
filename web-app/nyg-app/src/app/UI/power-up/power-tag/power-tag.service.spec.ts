import { TestBed } from '@angular/core/testing';

import { PowerTagService } from './power-tag.service';

describe('PowerTagService', () => {
  let service: PowerTagService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PowerTagService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
