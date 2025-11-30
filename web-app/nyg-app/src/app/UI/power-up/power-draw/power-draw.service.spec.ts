import { TestBed } from '@angular/core/testing';
import { DrawPowerService } from './power-draw.service';

describe('ViewGuessSheetService', () => {
  let service: DrawPowerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DrawPowerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
