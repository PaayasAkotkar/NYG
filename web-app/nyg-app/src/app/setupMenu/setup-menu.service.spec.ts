import { TestBed } from '@angular/core/testing';

import { SetupMenuService } from './setup-menu.service';

describe('SetupMenuService', () => {
  let service: SetupMenuService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SetupMenuService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
