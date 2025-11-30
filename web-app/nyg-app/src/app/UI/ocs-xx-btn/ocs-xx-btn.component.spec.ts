import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsXxBtnComponent } from './ocs-xx-btn.component';

describe('OcsXxBtnComponent', () => {
  let component: OcsXxBtnComponent;
  let fixture: ComponentFixture<OcsXxBtnComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsXxBtnComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsXxBtnComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
