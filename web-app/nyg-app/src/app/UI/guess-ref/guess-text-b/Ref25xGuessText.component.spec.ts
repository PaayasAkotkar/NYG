import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Ref25xGuessTextComponent } from './Ref25xGuessText.component';

describe('Ref25xGuessTextComponent ', () => {
  let component: Ref25xGuessTextComponent;
  let fixture: ComponentFixture<Ref25xGuessTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Ref25xGuessTextComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(Ref25xGuessTextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
