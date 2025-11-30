import { TestBed } from '@angular/core/testing';

import { OcsMiniviewService } from './ocs-miniview.service';

describe('OcsMiniviewService', () => {
  let service: OcsMiniviewService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OcsMiniviewService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
