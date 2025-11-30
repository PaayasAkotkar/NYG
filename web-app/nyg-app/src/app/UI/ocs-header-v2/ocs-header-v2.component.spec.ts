import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsHeaderV2Component } from './ocs-header-v2.component';

describe('OcsHeaderV2Component', () => {
  let component: OcsHeaderV2Component;
  let fixture: ComponentFixture<OcsHeaderV2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsHeaderV2Component]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsHeaderV2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
