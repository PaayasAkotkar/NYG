import { TestBed } from '@angular/core/testing';
import { NexusPowerService } from './power-nexus.service';

describe('ViewGuessSheetService', () => {
  let service: NexusPowerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NexusPowerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
