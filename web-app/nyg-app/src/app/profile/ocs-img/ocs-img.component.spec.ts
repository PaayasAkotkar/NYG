import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsImgComponent } from './ocs-img.component';

describe('OcsImgComponent', () => {
  let component: OcsImgComponent;
  let fixture: ComponentFixture<OcsImgComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsImgComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsImgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
