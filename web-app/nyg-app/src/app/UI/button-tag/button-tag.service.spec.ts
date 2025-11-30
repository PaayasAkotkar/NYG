import { TestBed } from '@angular/core/testing';

import { ButtonTagService } from './button-tag.service';

describe('ButtonTagService', () => {
  let service: ButtonTagService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ButtonTagService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
