import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsInputComponent } from './ocs-input.component';

describe('OcsInputComponent', () => {
  let component: OcsInputComponent;
  let fixture: ComponentFixture<OcsInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsInputComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
