import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GuessText11xComponent } from './guess-text-11x.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('GuessText11xComponent', () => {
  let component: GuessText11xComponent;
  let fixture: ComponentFixture<GuessText11xComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, GuessText11xComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(GuessText11xComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
