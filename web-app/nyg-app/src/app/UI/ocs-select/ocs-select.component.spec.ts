import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsSelectComponent } from './ocs-select.component';

describe('OcsSelectComponent', () => {
  let component: OcsSelectComponent;
  let fixture: ComponentFixture<OcsSelectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsSelectComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsSelectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
