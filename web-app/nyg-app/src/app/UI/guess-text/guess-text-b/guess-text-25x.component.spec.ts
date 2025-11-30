import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GuessText25xComponent } from './guess-text-25x.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('GuessText25xComponent', () => {
  let component: GuessText25xComponent;
  let fixture: ComponentFixture<GuessText25xComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, GuessText25xComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(GuessText25xComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
