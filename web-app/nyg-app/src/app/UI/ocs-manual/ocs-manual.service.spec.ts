import { TestBed } from '@angular/core/testing';

import { OcsManualService } from './ocs-manual.service';

describe('OcsManualService', () => {
  let service: OcsManualService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OcsManualService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
