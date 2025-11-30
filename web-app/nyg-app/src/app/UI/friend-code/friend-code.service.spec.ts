import { TestBed } from '@angular/core/testing';
import { FriendCodeService } from './friend-code.service';
describe('ViewGuessSheetService', () => {
  let service: FriendCodeService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FriendCodeService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
