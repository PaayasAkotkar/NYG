import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PickPatternV2Component } from './pick-pattern-v2.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('PickPatternV2Component', () => {
  let component: PickPatternV2Component;
  let fixture: ComponentFixture<PickPatternV2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, PickPatternV2Component]
    })
      .compileComponents();

    fixture = TestBed.createComponent(PickPatternV2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
