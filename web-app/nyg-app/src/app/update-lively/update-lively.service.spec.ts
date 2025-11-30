import { TestBed } from '@angular/core/testing';

import { UpdateLivelyService } from './update-lively.service';
import { HttpClient, provideHttpClient } from '@angular/common/http';
import { HttpClientTestingModule, provideHttpClientTesting } from '@angular/common/http/testing';
describe('UpdateLivelyService', () => {
  let service: UpdateLivelyService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [HttpClient, provideHttpClient(), provideHttpClientTesting()]
    });

    service = TestBed.inject(UpdateLivelyService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
