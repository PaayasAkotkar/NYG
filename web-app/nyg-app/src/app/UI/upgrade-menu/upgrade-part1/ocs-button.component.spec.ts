import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OCSbuttonV1 } from './ocs-button.component';

describe('OCSbuttonV1', () => {
  let component: OCSbuttonV1;
  let fixture: ComponentFixture<OCSbuttonV1>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OCSbuttonV1]
    })
      .compileComponents();

    fixture = TestBed.createComponent(OCSbuttonV1);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
