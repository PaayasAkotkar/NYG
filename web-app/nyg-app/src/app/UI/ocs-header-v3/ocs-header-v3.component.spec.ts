import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsHeaderV3Component } from './ocs-header-v3.component';

describe('OcsHeaderV3Component', () => {
  let component: OcsHeaderV3Component;
  let fixture: ComponentFixture<OcsHeaderV3Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsHeaderV3Component]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsHeaderV3Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
