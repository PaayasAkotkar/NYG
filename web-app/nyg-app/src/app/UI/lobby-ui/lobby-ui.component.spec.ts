import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LobbyUiComponent } from './lobby-ui.component';

describe('LobbyUiComponent', () => {
  let component: LobbyUiComponent;
  let fixture: ComponentFixture<LobbyUiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LobbyUiComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LobbyUiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
