import { TestBed } from '@angular/core/testing';

import { PartyChatService } from './party-chat.service';

describe('PartyChatService', () => {
  let service: PartyChatService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PartyChatService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
