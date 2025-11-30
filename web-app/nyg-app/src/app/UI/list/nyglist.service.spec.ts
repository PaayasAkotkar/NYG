import { TestBed } from '@angular/core/testing';

import { NYGlistService } from './nyglist.service';

describe('NYGlistService', () => {
  let service: NYGlistService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NYGlistService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
