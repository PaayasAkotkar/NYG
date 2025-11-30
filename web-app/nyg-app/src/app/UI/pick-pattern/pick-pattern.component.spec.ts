import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PickPatternComponent } from './pick-pattern.component';
import { FormControl, FormGroup } from '@angular/forms';

describe('PickPatternComponent', () => {
  let component: PickPatternComponent;
  let fixture: ComponentFixture<PickPatternComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PickPatternComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(PickPatternComponent);
    component = fixture.componentInstance;
    component.OCSformGroup = new FormGroup({
      N: new FormControl("")
    })
    fixture.detectChanges();
  });
  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
