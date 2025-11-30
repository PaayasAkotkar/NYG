import { TestBed } from '@angular/core/testing';
import { BetPowerService } from './power-bet.service';
import { BetGuessComponent } from './bet-picker/bet-guess.component';
import { beforeEach } from 'node:test';

describe('BetPowerService', () => {
  let service: BetPowerService;
  afterEach(() => {
    TestBed.resetTestingModule();
  });
  beforeEach(async () => {
    TestBed.configureTestingModule({
      providers: [{ provide: BetGuessComponent, useValue: {} }]
    });
    service = TestBed.inject(BetPowerService);
  })
  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
