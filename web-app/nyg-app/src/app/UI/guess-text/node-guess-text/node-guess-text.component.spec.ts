import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NodeGuessTextComponent } from './node-guess-text.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('NodeGuessTextComponent', () => {
  let component: NodeGuessTextComponent;
  let fixture: ComponentFixture<NodeGuessTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, NodeGuessTextComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(NodeGuessTextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
