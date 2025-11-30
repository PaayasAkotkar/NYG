import { TestBed } from '@angular/core/testing';
import { FrameService } from './frame.service';

describe('FrameSerivce', () => {
  let service: FrameService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FrameService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
