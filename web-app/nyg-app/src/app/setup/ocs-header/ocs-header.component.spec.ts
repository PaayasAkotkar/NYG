import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsHeaderComponent } from './ocs-header.component';

describe('OcsHeaderComponent', () => {
  let component: OcsHeaderComponent;
  let fixture: ComponentFixture<OcsHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsHeaderComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
