import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GameTicketsComponent } from './game-tickets.component';

describe('GameTicketsComponent', () => {
  let component: GameTicketsComponent;
  let fixture: ComponentFixture<GameTicketsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GameTicketsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GameTicketsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
