import { TestBed } from '@angular/core/testing';
import { FreezePowerService } from './power-freeze.service';

describe('ViewGuessSheetService', () => {
  let service: FreezePowerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FreezePowerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
