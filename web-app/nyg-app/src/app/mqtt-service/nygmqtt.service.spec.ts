import { TestBed } from '@angular/core/testing';

import { NYGmqttService } from './nygmqtt.service';
import { HttpClient } from '@angular/common/http';

describe('NYGmqttService', () => {
  let service: NYGmqttService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [HttpClient]
    });
    service = TestBed.inject(NYGmqttService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
