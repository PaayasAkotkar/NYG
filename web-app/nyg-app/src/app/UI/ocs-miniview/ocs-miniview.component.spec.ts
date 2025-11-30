import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsMiniviewComponent } from './ocs-miniview.component';

describe('OcsMiniviewComponent', () => {
  let component: OcsMiniviewComponent;
  let fixture: ComponentFixture<OcsMiniviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsMiniviewComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsMiniviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
