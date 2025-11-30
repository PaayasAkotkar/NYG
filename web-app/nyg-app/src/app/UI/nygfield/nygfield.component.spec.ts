import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NYGfieldV2Component } from './nygfield.component';

describe('NYGfieldV2Component', () => {
  let component: NYGfieldV2Component;
  let fixture: ComponentFixture<NYGfieldV2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NYGfieldV2Component]
    })
      .compileComponents();

    fixture = TestBed.createComponent(NYGfieldV2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
