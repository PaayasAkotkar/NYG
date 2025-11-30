import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Ref11xGuessTextComponent } from './Ref11xGuessText.component';

describe('Ref11xGuessTextComponent ', () => {
  let component: Ref11xGuessTextComponent;
  let fixture: ComponentFixture<Ref11xGuessTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Ref11xGuessTextComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(Ref11xGuessTextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
