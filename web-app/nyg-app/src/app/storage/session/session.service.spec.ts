import { TestBed } from '@angular/core/testing';

import { SessionService } from './session.service';

describe('SessionService', () => {
  let service: SessionService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SessionService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
  it('should set the value', () => {
    let testValue = '42342'
    spyOn(sessionStorage, 'getItem').and.returnValue(testValue)

    const result = service.getItem(testValue);

    expect(sessionStorage.getItem).toHaveBeenCalledWith(testValue);
    expect(result).toBe(testValue);
  })
});
