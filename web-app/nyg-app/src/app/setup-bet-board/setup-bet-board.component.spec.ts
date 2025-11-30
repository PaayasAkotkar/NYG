import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetupBetBoardComponent } from './setup-bet-board.component';

describe('SetupBetBoardComponent', () => {
  let component: SetupBetBoardComponent;
  let fixture: ComponentFixture<SetupBetBoardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetupBetBoardComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SetupBetBoardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
