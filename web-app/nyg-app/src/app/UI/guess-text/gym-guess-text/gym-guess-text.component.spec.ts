import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GymGuessTextComponent } from './gym-guess-text.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('GymGuessTextComponent', () => {
  let component: GymGuessTextComponent;
  let fixture: ComponentFixture<GymGuessTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, GymGuessTextComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(GymGuessTextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
