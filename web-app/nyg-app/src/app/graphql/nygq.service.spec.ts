import { TestBed } from '@angular/core/testing';

import { NYGqService } from './nygq.service';
import { ApolloTestingController, ApolloTestingModule } from 'apollo-angular/testing';
import { Apollo } from 'apollo-angular';
describe('NYGqService', () => {
  let service: NYGqService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ApolloTestingModule],
      providers: [Apollo]
    });
    service = TestBed.inject(NYGqService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
