import { TestBed } from '@angular/core/testing';
import { EntertainmentService } from './setup-entertainment.service';

describe('SetupMenuService', () => {
  let service: EntertainmentService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(EntertainmentService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
