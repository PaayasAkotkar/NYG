import { TestBed } from '@angular/core/testing';

import { OnLoadService } from './on-load.service';

describe('OnLoadService', () => {
  let service: OnLoadService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OnLoadService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
