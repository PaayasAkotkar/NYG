import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OCSFrameComponent } from './frame.component';

describe('LeftOverComponent', () => {
  let component: OCSFrameComponent;
  let fixture: ComponentFixture<OCSFrameComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OCSFrameComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(OCSFrameComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
