import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsBtn2Component } from './ocs-btn-2.component';

describe('OcsBtn2Component', () => {
  let component: OcsBtn2Component;
  let fixture: ComponentFixture<OcsBtn2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsBtn2Component]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsBtn2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
