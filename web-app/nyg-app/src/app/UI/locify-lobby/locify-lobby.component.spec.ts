import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LocifyLobbyComponent } from './locify-lobby.component';

describe('LocifyLobbyComponent', () => {
  let component: LocifyLobbyComponent;
  let fixture: ComponentFixture<LocifyLobbyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LocifyLobbyComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LocifyLobbyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
