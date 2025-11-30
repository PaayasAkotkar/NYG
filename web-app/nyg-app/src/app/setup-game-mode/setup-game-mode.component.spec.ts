import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetupGameModeComponent } from './setup-game-mode.component';

describe('SetupGameModeComponent', () => {
  let component: SetupGameModeComponent;
  let fixture: ComponentFixture<SetupGameModeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetupGameModeComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(SetupGameModeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
