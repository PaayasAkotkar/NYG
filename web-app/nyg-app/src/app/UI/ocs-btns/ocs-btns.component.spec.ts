import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsBtnsComponent } from './ocs-btns.component';

describe('OcsBtnsComponent', () => {
  let component: OcsBtnsComponent;
  let fixture: ComponentFixture<OcsBtnsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsBtnsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsBtnsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
