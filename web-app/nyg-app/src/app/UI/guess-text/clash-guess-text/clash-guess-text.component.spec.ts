import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClashGuessTextComponent } from './clash-guess-text.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('ClashGuessTextComponent', () => {
  let component: ClashGuessTextComponent;
  let fixture: ComponentFixture<ClashGuessTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, ClashGuessTextComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(ClashGuessTextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
