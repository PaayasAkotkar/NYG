import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BetGuessComponent } from './bet-guess.component';
import { BetPowerService } from '../power-bet.service';
import { ReactiveFormsModule } from '@angular/forms';

describe('BetGuessComponent', () => {
  let component: BetGuessComponent;
  let fixture: ComponentFixture<BetGuessComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BetGuessComponent, ReactiveFormsModule],
      providers: [BetPowerService]
    })
      .compileComponents();

    fixture = TestBed.createComponent(BetGuessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
