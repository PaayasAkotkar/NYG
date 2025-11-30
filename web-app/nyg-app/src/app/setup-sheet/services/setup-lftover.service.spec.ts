import { TestBed } from '@angular/core/testing';
import { LftOverService } from './setup-lftover.service';

describe('SetupMenuService', () => {
  let service: LftOverService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LftOverService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
