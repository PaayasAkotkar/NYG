import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GuessTextComponent } from './guess-text.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('GuessTextComponent', () => {
  let component: GuessTextComponent;
  let fixture: ComponentFixture<GuessTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, GuessTextComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(GuessTextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
