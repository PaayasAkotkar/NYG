import { TestBed } from '@angular/core/testing';

import { SpotLightService } from './spot-light.service';

describe('SpotLightService', () => {
  let service: SpotLightService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SpotLightService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
