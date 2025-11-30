import { TestBed } from '@angular/core/testing';

import { HomePagePorfilesService } from './home-page-porfiles.service';

describe('HomePagePorfilesService', () => {
  let service: HomePagePorfilesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HomePagePorfilesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
