import { TestBed } from '@angular/core/testing';

import { OcsIndoxMsgsService } from './ocs-indox-msgs.service';

describe('OcsIndoxMsgsService', () => {
  let service: OcsIndoxMsgsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OcsIndoxMsgsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
