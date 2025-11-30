import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GuessSheetComponent } from './guess-sheet.component';

describe('GuessSheetComponent', () => {
  let component: GuessSheetComponent;
  let fixture: ComponentFixture<GuessSheetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GuessSheetComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GuessSheetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
