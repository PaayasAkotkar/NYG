import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsImageComponent } from './ocs-image.component';

describe('OcsImageComponent', () => {
  let component: OcsImageComponent;
  let fixture: ComponentFixture<OcsImageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsImageComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsImageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
