import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetupBetGuessComponent } from './setup-bet-guess.component';

describe('SetupBetGuessComponent', () => {
  let component: SetupBetGuessComponent;
  let fixture: ComponentFixture<SetupBetGuessComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetupBetGuessComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SetupBetGuessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
