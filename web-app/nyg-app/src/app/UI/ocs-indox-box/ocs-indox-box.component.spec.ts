import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsIndoxBoxComponent } from './ocs-indox-box.component';

describe('OcsIndoxBoxComponent', () => {
  let component: OcsIndoxBoxComponent;
  let fixture: ComponentFixture<OcsIndoxBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsIndoxBoxComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsIndoxBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
