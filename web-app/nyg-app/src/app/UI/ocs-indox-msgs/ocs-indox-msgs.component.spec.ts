import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsIndoxMsgsComponent } from './ocs-indox-msgs.component';

describe('OcsIndoxMsgsComponent', () => {
  let component: OcsIndoxMsgsComponent;
  let fixture: ComponentFixture<OcsIndoxMsgsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsIndoxMsgsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsIndoxMsgsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
