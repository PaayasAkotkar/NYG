import { TestBed } from '@angular/core/testing';

import { OcsInputServiceV2 } from './ocs-input-service-v2.service';

describe('OcsInputServiceV2Service', () => {
  let service: OcsInputServiceV2;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OcsInputServiceV2);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
