import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsSaveImgComponent } from './ocs-save-img.component';

describe('OcsSaveImgComponent', () => {
  let component: OcsSaveImgComponent;
  let fixture: ComponentFixture<OcsSaveImgComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsSaveImgComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsSaveImgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
