import { TestBed } from '@angular/core/testing';
import { BetGuessComponent } from './bet-guess.component';

describe('BetPowerService', () => {
  let service: BetGuessComponent;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BetGuessComponent);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
