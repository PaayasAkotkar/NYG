import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsCounterComponent } from './ocs-counter.component';

describe('OcsCounterComponent', () => {
  let component: OcsCounterComponent;
  let fixture: ComponentFixture<OcsCounterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsCounterComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsCounterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
