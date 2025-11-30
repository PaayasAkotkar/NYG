import { TestBed } from '@angular/core/testing';
import { RoomSettingsComponent } from './room-settings.component';
describe('ViewGuessSheetService', () => {
  let service: RoomSettingsComponent;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RoomSettingsComponent);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
