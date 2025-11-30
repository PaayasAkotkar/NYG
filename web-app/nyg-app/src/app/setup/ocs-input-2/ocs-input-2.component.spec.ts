import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsInput2Component } from './ocs-input-2.component';

describe('OcsInput2Component', () => {
  let component: OcsInput2Component;
  let fixture: ComponentFixture<OcsInput2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsInput2Component]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsInput2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
