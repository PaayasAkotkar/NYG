import { TestBed } from '@angular/core/testing';

import { LoadLogoService } from './load-logo.service';

describe('LoadLogoService', () => {
  let service: LoadLogoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LoadLogoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
