import { TestBed } from '@angular/core/testing';

import { OcsInputService } from './ocs-input.service';

describe('OcsInputService', () => {
  let service: OcsInputService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OcsInputService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
