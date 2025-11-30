import { TestBed } from '@angular/core/testing';

import { ViewGuessSheetService } from './view-guess-sheet.service';

describe('ViewGuessSheetService', () => {
  let service: ViewGuessSheetService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ViewGuessSheetService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
